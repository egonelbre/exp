package deferslice

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "deferslice",
	Doc:      "reports deferred calls where a slice argument is modified later in the function",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		var body *ast.BlockStmt
		switch fn := n.(type) {
		case *ast.FuncDecl:
			body = fn.Body
		case *ast.FuncLit:
			body = fn.Body
		}
		if body == nil {
			return
		}
		checkFunc(pass, body)
	})

	return nil, nil
}

// deferredSlice records a slice variable passed to a defer call.
type deferredSlice struct {
	pos  token.Pos    // position of the defer statement
	obj  types.Object // the slice variable
	name string       // variable name for diagnostics
}

func checkFunc(pass *analysis.Pass, body *ast.BlockStmt) {
	// Collect all defer statements with slice arguments.
	var deferred []deferredSlice
	inspectSkippingFuncLit(body, func(n ast.Node) {
		ds, ok := n.(*ast.DeferStmt)
		if !ok {
			return
		}
		for i, arg := range ds.Call.Args {
			ident, ok := arg.(*ast.Ident)
			if !ok {
				continue
			}
			obj := pass.TypesInfo.ObjectOf(ident)
			if obj == nil {
				continue
			}
			if _, ok := obj.Type().Underlying().(*types.Slice); ok {
				if isRestorePattern(pass, ds, i, obj) {
					continue
				}
				deferred = append(deferred, deferredSlice{
					pos:  ds.Pos(),
					obj:  obj,
					name: ident.Name,
				})
			}
		}
	})

	if len(deferred) == 0 {
		return
	}

	// Build a set of deferred slice objects for quick lookup.
	deferredByObj := map[types.Object][]deferredSlice{}
	for _, d := range deferred {
		deferredByObj[d.obj] = append(deferredByObj[d.obj], d)
	}

	// Find modifications to deferred slices that occur after the defer.
	inspectSkippingFuncLit(body, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			for _, lhs := range n.Lhs {
				obj, name := sliceModTarget(pass, lhs)
				if obj == nil {
					continue
				}
				for _, d := range deferredByObj[obj] {
					if d.pos < n.Pos() {
						pass.Reportf(n.Pos(), "slice %s is modified after being passed to defer", name)
						break
					}
				}
			}
		case *ast.CallExpr:
			name := calleeName(n)
			isMutating := mutatingFuncs[name]
			for _, arg := range n.Args {
				// Check for &slice (pointer to slice implies mutation).
				if unary, ok := arg.(*ast.UnaryExpr); ok && unary.Op == token.AND {
					if ident, ok := unary.X.(*ast.Ident); ok {
						obj := pass.TypesInfo.ObjectOf(ident)
						if obj == nil {
							continue
						}
						for _, d := range deferredByObj[obj] {
							if d.pos < n.Pos() {
								pass.Reportf(n.Pos(), "pointer to slice %s is passed to %s after being passed to defer", ident.Name, name)
								break
							}
						}
					}
					continue
				}
				// Check for known mutating functions.
				if !isMutating {
					continue
				}
				ident, ok := arg.(*ast.Ident)
				if !ok {
					continue
				}
				obj := pass.TypesInfo.ObjectOf(ident)
				if obj == nil {
					continue
				}
				for _, d := range deferredByObj[obj] {
					if d.pos < n.Pos() {
						pass.Reportf(n.Pos(), "slice %s is passed to %s after being passed to defer", ident.Name, name)
						break
					}
				}
			}
		}
	})
}

// mutatingFuncs is the set of function/method names that are known to modify
// a slice argument in place.
var mutatingFuncs = map[string]bool{
	"Read":        true,
	"ReadAt":      true,
	"ReadFull":    true,
	"ReadAtLeast": true,
	"copy":        true,
}

// calleeName returns the simple name of the called function or method.
func calleeName(call *ast.CallExpr) string {
	switch fn := call.Fun.(type) {
	case *ast.Ident:
		return fn.Name
	case *ast.SelectorExpr:
		return fn.Sel.Name
	}
	return ""
}

// sliceModTarget returns the types.Object and name if expr is a modification
// of a slice variable: either `s[i]` or `s` (direct reassignment).
func sliceModTarget(pass *analysis.Pass, expr ast.Expr) (types.Object, string) {
	switch e := expr.(type) {
	case *ast.IndexExpr:
		if ident, ok := e.X.(*ast.Ident); ok {
			obj := pass.TypesInfo.ObjectOf(ident)
			if obj != nil {
				if _, ok := obj.Type().Underlying().(*types.Slice); ok {
					return obj, ident.Name
				}
			}
		}
	case *ast.Ident:
		obj := pass.TypesInfo.ObjectOf(e)
		if obj != nil {
			if _, ok := obj.Type().Underlying().(*types.Slice); ok {
				return obj, e.Name
			}
		}
	}
	return nil, ""
}

// isRestorePattern detects the save/restore idiom:
//
//	defer func(old []T) { x = old }(x)
//
// where the defer captures the current value and restores it on return.
func isRestorePattern(pass *analysis.Pass, ds *ast.DeferStmt, argIndex int, argObj types.Object) bool {
	funcLit, ok := ds.Call.Fun.(*ast.FuncLit)
	if !ok {
		return false
	}

	// Find the parameter name at the given argument index.
	paramIdent := paramAt(funcLit.Type.Params, argIndex)
	if paramIdent == nil {
		return false
	}
	paramObj := pass.TypesInfo.ObjectOf(paramIdent)
	if paramObj == nil {
		return false
	}

	// Check whether the body assigns the parameter back to the original variable.
	for _, stmt := range funcLit.Body.List {
		assign, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}
		for i, lhs := range assign.Lhs {
			lhsIdent, ok := lhs.(*ast.Ident)
			if !ok {
				continue
			}
			if i >= len(assign.Rhs) {
				continue
			}
			rhsIdent, ok := assign.Rhs[i].(*ast.Ident)
			if !ok {
				continue
			}
			if pass.TypesInfo.ObjectOf(lhsIdent) == argObj &&
				pass.TypesInfo.ObjectOf(rhsIdent) == paramObj {
				return true
			}
		}
	}
	return false
}

// paramAt returns the *ast.Ident of the parameter at the given index,
// accounting for grouped parameters (e.g. func(a, b []int)).
func paramAt(params *ast.FieldList, index int) *ast.Ident {
	i := 0
	for _, field := range params.List {
		for _, name := range field.Names {
			if i == index {
				return name
			}
			i++
		}
	}
	return nil
}

// inspectSkippingFuncLit walks the AST calling fn for each node,
// but does not descend into nested function literals.
func inspectSkippingFuncLit(node ast.Node, fn func(ast.Node)) {
	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		if _, ok := n.(*ast.FuncLit); ok {
			return false
		}
		fn(n)
		return true
	})
}
