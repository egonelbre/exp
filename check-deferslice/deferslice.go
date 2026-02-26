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
		for _, arg := range ds.Call.Args {
			ident, ok := arg.(*ast.Ident)
			if !ok {
				continue
			}
			obj := pass.TypesInfo.ObjectOf(ident)
			if obj == nil {
				continue
			}
			if _, ok := obj.Type().Underlying().(*types.Slice); ok {
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
			if !mutatingFuncs[name] {
				return
			}
			for _, arg := range n.Args {
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
