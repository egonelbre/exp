package untangle

import (
	"go/ast"
	"go/types"
	"sort"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:      "untangle",
	Doc:       "tracks field access patterns for structs marked with //untangle:track",
	Run:       run,
	FactTypes: []analysis.Fact{&TrackedStructFact{}, &AccessPatternFact{}},
}

// TrackedStructFact records that a struct type is being tracked
type TrackedStructFact struct {
	Fields []string // field names in the tracked struct
}

func (f *TrackedStructFact) AFact() {}

func (f *TrackedStructFact) String() string {
	return "tracked struct with fields: " + strings.Join(f.Fields, ", ")
}

// AccessPatternFact records which fields of tracked structs are accessed by a function/method
type AccessPatternFact struct {
	// Map from struct type name to set of field names accessed
	Accesses map[string]map[string]bool
}

func (f *AccessPatternFact) AFact() {}

func (f *AccessPatternFact) String() string {
	if len(f.Accesses) == 0 {
		return "no accesses"
	}
	var parts []string
	for typeName, fields := range f.Accesses {
		fieldList := make([]string, 0, len(fields))
		for field := range fields {
			fieldList = append(fieldList, field)
		}
		sort.Strings(fieldList)
		parts = append(parts, typeName+": "+strings.Join(fieldList, ", "))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func run(pass *analysis.Pass) (interface{}, error) {
	// Stage 1: Find tracked structs
	trackedStructs := findTrackedStructs(pass)

	// Stage 2: Analyze field accesses for all functions/methods
	analyzeFunctions(pass, trackedStructs)

	// Stage 3: Report for functions marked with //untangle:print
	reportAccessPatterns(pass)

	return nil, nil
}

// Stage 1: Find all structs marked with //untangle:track
func findTrackedStructs(pass *analysis.Pass) map[*types.TypeName]bool {
	tracked := make(map[*types.TypeName]bool)

	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				// Check for //untangle:track comment
				if !hasDirective(genDecl.Doc, "untangle:track") {
					continue
				}

				// Get the TypeName object
				obj := pass.TypesInfo.Defs[typeSpec.Name]
				typeName, ok := obj.(*types.TypeName)
				if !ok {
					continue
				}

				// Extract field names
				var fields []string
				for _, field := range structType.Fields.List {
					for _, name := range field.Names {
						fields = append(fields, name.Name)
					}
				}

				// Export fact
				fact := &TrackedStructFact{Fields: fields}
				pass.ExportObjectFact(typeName, fact)
				tracked[typeName] = true
			}
		}
	}

	return tracked
}

// Stage 2: Analyze field accesses for all functions and methods
func analyzeFunctions(pass *analysis.Pass, localTracked map[*types.TypeName]bool) {
	// Step 1: Collect all functions to analyze
	funcs := collectFunctions(pass)

	// Step 2: Build call graph and find SCCs (strongly connected components)
	callGraph := buildCallGraph(pass, funcs)
	sccs := findSCCs(callGraph, funcs)

	// Step 3: Topologically sort SCCs
	sortedSCCs := topologicalSort(sccs, callGraph)

	// Step 4: Process functions in topological order
	for _, scc := range sortedSCCs {
		// Process all functions in this SCC together (iteratively for recursion)
		processSCC(pass, scc, localTracked)
	}
}

// collectFunctions gathers all function declarations that should be analyzed
func collectFunctions(pass *analysis.Pass) []*funcInfo {
	var funcs []*funcInfo

	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}

			// Check for //untangle:ignore
			if hasDirective(funcDecl.Doc, "untangle:ignore") {
				continue
			}

			// Get the function object
			obj := pass.TypesInfo.Defs[funcDecl.Name]
			funcObj, ok := obj.(*types.Func)
			if !ok {
				continue
			}

			funcs = append(funcs, &funcInfo{
				decl: funcDecl,
				obj:  funcObj,
			})
		}
	}

	return funcs
}

// funcInfo holds information about a function to analyze
type funcInfo struct {
	decl *ast.FuncDecl
	obj  *types.Func
}

// buildCallGraph constructs a call graph among the functions
func buildCallGraph(pass *analysis.Pass, funcs []*funcInfo) map[*types.Func][]*types.Func {
	callGraph := make(map[*types.Func][]*types.Func)
	funcSet := make(map[*types.Func]bool)

	// Build a set of all functions we're tracking
	for _, f := range funcs {
		funcSet[f.obj] = true
		callGraph[f.obj] = nil // Initialize entry
	}

	// For each function, find which other tracked functions it calls
	for _, f := range funcs {
		if f.decl.Body == nil {
			continue
		}

		callees := make(map[*types.Func]bool)
		ast.Inspect(f.decl.Body, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			callee := getCalledFunc(pass, call)
			if callee != nil && funcSet[callee] {
				callees[callee] = true
			}
			return true
		})

		// Convert set to slice
		for callee := range callees {
			callGraph[f.obj] = append(callGraph[f.obj], callee)
		}
	}

	return callGraph
}

// getCalledFunc extracts the *types.Func being called from a CallExpr
func getCalledFunc(pass *analysis.Pass, call *ast.CallExpr) *types.Func {
	switch fun := call.Fun.(type) {
	case *ast.SelectorExpr:
		sel := pass.TypesInfo.Selections[fun]
		if sel != nil && sel.Kind() == types.MethodVal {
			funcObj, _ := sel.Obj().(*types.Func)
			return funcObj
		}
		if obj := pass.TypesInfo.Uses[fun.Sel]; obj != nil {
			funcObj, _ := obj.(*types.Func)
			return funcObj
		}

	case *ast.Ident:
		if obj := pass.TypesInfo.Uses[fun]; obj != nil {
			funcObj, _ := obj.(*types.Func)
			return funcObj
		}
	}
	return nil
}

// findSCCs uses Tarjan's algorithm to find strongly connected components
func findSCCs(callGraph map[*types.Func][]*types.Func, funcs []*funcInfo) [][]*types.Func {
	index := 0
	stack := []*types.Func{}
	indices := make(map[*types.Func]int)
	lowlinks := make(map[*types.Func]int)
	onStack := make(map[*types.Func]bool)
	var sccs [][]*types.Func

	var strongConnect func(*types.Func)
	strongConnect = func(v *types.Func) {
		indices[v] = index
		lowlinks[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		for _, w := range callGraph[v] {
			if _, visited := indices[w]; !visited {
				strongConnect(w)
				if lowlinks[w] < lowlinks[v] {
					lowlinks[v] = lowlinks[w]
				}
			} else if onStack[w] {
				if indices[w] < lowlinks[v] {
					lowlinks[v] = indices[w]
				}
			}
		}

		if lowlinks[v] == indices[v] {
			var scc []*types.Func
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			sccs = append(sccs, scc)
		}
	}

	for _, f := range funcs {
		if _, visited := indices[f.obj]; !visited {
			strongConnect(f.obj)
		}
	}

	return sccs
}

// topologicalSort orders SCCs based on dependencies
func topologicalSort(sccs [][]*types.Func, callGraph map[*types.Func][]*types.Func) [][]*types.Func {
	// Build SCC index map
	sccIndex := make(map[*types.Func]int)
	for i, scc := range sccs {
		for _, f := range scc {
			sccIndex[f] = i
		}
	}

	// Build SCC dependency graph
	sccDeps := make(map[int]map[int]bool)
	for i := range sccs {
		sccDeps[i] = make(map[int]bool)
	}

	for i, scc := range sccs {
		for _, f := range scc {
			for _, callee := range callGraph[f] {
				calleeIdx := sccIndex[callee]
				if calleeIdx != i {
					sccDeps[i][calleeIdx] = true
				}
			}
		}
	}

	// Topological sort using DFS
	visited := make(map[int]bool)
	var result [][]*types.Func

	var visit func(int)
	visit = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true

		for dep := range sccDeps[i] {
			visit(dep)
		}

		result = append(result, sccs[i])
	}

	for i := range sccs {
		visit(i)
	}

	return result
}

// processSCC analyzes all functions in an SCC together
func processSCC(pass *analysis.Pass, scc []*types.Func, localTracked map[*types.TypeName]bool) {
	// Build a map from function object to declaration
	funcDecls := make(map[*types.Func]*ast.FuncDecl)
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			obj := pass.TypesInfo.Defs[funcDecl.Name]
			if funcObj, ok := obj.(*types.Func); ok {
				funcDecls[funcObj] = funcDecl
			}
		}
	}

	// For SCCs with multiple functions (recursive), iterate until stable
	maxIterations := 10
	for iter := 0; iter < maxIterations; iter++ {
		changed := false

		for _, funcObj := range scc {
			funcDecl := funcDecls[funcObj]
			if funcDecl == nil {
				continue
			}

			// Analyze field accesses in this function
			accesses := analyzeFieldAccesses(pass, funcDecl, localTracked)

			// Check if this differs from existing fact
			var oldFact AccessPatternFact
			hasOldFact := pass.ImportObjectFact(funcObj, &oldFact)

			if !hasOldFact || !accessesEqual(accesses, oldFact.Accesses) {
				changed = true
				if len(accesses) > 0 {
					fact := &AccessPatternFact{Accesses: accesses}
					pass.ExportObjectFact(funcObj, fact)
				}
			}
		}

		// If no changes in this iteration, we've reached a fixed point
		if !changed {
			break
		}
	}
}

// accessesEqual checks if two access maps are equal
func accessesEqual(a, b map[string]map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for typeName, fieldsA := range a {
		fieldsB, ok := b[typeName]
		if !ok || len(fieldsA) != len(fieldsB) {
			return false
		}
		for field := range fieldsA {
			if !fieldsB[field] {
				return false
			}
		}
	}
	return true
}

// analyzeFieldAccesses finds all field accesses of tracked structs in a function
func analyzeFieldAccesses(pass *analysis.Pass, funcDecl *ast.FuncDecl, localTracked map[*types.TypeName]bool) map[string]map[string]bool {
	accesses := make(map[string]map[string]bool)

	// Skip functions without a body (declarations, interface methods, etc.)
	if funcDecl.Body == nil {
		return accesses
	}

	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.CallExpr:
			// Check if this is a method call or function call that accesses tracked fields
			// Process CallExpr BEFORE SelectorExpr to avoid double-processing
			handleFunctionCall(pass, node, localTracked, accesses)

		case *ast.SelectorExpr:
			// Check if this is a field access on a tracked struct
			handleFieldAccess(pass, node, localTracked, accesses)
		}
		return true
	})

	return accesses
}

// handleFieldAccess processes a selector expression (obj.Field)
func handleFieldAccess(pass *analysis.Pass, sel *ast.SelectorExpr, localTracked map[*types.TypeName]bool, accesses map[string]map[string]bool) {
	// Check if this is actually a field selection (not a method)
	selection := pass.TypesInfo.Selections[sel]
	if selection == nil {
		// Package-qualified identifier or other non-selection
		return
	}
	if selection.Kind() != types.FieldVal {
		// Method call, not a field access
		return
	}

	// Get the type of the base expression
	baseType := pass.TypesInfo.TypeOf(sel.X)
	if baseType == nil {
		return
	}

	// Strip pointer if present
	if ptr, ok := baseType.(*types.Pointer); ok {
		baseType = ptr.Elem()
	}

	// Check if it's a named type
	named, ok := baseType.(*types.Named)
	if !ok {
		return
	}

	// Check if this type is tracked
	typeName := named.Obj()
	if !isTrackedStruct(pass, typeName, localTracked) {
		return
	}

	// Record the field access
	fieldName := sel.Sel.Name
	typeNameStr := typeName.Pkg().Path() + "." + typeName.Name()

	if accesses[typeNameStr] == nil {
		accesses[typeNameStr] = make(map[string]bool)
	}
	accesses[typeNameStr][fieldName] = true
}

// handleFunctionCall processes function and method calls to propagate access patterns
func handleFunctionCall(pass *analysis.Pass, call *ast.CallExpr, localTracked map[*types.TypeName]bool, accesses map[string]map[string]bool) {
	var calleeObj *types.Func

	switch fun := call.Fun.(type) {
	case *ast.SelectorExpr:
		// Method call or qualified function call
		sel := pass.TypesInfo.Selections[fun]
		if sel != nil && sel.Kind() == types.MethodVal {
			// Method call
			calleeObj, _ = sel.Obj().(*types.Func)
		} else {
			// Qualified function call
			if obj := pass.TypesInfo.Uses[fun.Sel]; obj != nil {
				calleeObj, _ = obj.(*types.Func)
			}
		}

	case *ast.Ident:
		// Direct function call
		if obj := pass.TypesInfo.Uses[fun]; obj != nil {
			calleeObj, _ = obj.(*types.Func)
		}
	}

	if calleeObj == nil {
		return
	}

	// Import the callee's access pattern fact
	var calleeFact AccessPatternFact
	if pass.ImportObjectFact(calleeObj, &calleeFact) {
		// Merge the callee's accesses into our accesses
		mergeAccesses(accesses, calleeFact.Accesses)
	}
}

// mergeAccesses merges source accesses into destination
func mergeAccesses(dest, src map[string]map[string]bool) {
	for typeName, fields := range src {
		if dest[typeName] == nil {
			dest[typeName] = make(map[string]bool)
		}
		for field := range fields {
			dest[typeName][field] = true
		}
	}
}

// isTrackedStruct checks if a type name refers to a tracked struct
func isTrackedStruct(pass *analysis.Pass, typeName *types.TypeName, localTracked map[*types.TypeName]bool) bool {
	// Check local tracked structs
	if localTracked[typeName] {
		return true
	}

	// Check imported facts
	var fact TrackedStructFact
	return pass.ImportObjectFact(typeName, &fact)
}

// Stage 3: Report access patterns for functions marked with //untangle:print
func reportAccessPatterns(pass *analysis.Pass) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}

			// Check for //untangle:print
			if !hasDirective(funcDecl.Doc, "untangle:print") {
				continue
			}

			// Get the function object
			obj := pass.TypesInfo.Defs[funcDecl.Name]
			funcObj, ok := obj.(*types.Func)
			if !ok {
				continue
			}

			// Import the function's access pattern fact
			var fact AccessPatternFact
			if pass.ImportObjectFact(funcObj, &fact) {
				// Format and report the access pattern
				msg := formatAccessPattern(funcObj.Name(), fact.Accesses)
				pass.Reportf(funcDecl.Pos(), "%s", msg)
			} else {
				// No accesses recorded
				pass.Reportf(funcDecl.Pos(), "%s: no tracked field accesses", funcObj.Name())
			}
		}
	}
}

// formatAccessPattern creates a human-readable access pattern report
func formatAccessPattern(funcName string, accesses map[string]map[string]bool) string {
	if len(accesses) == 0 {
		return funcName + ": no tracked field accesses"
	}

	var parts []string
	for typeName, fields := range accesses {
		fieldList := make([]string, 0, len(fields))
		for field := range fields {
			fieldList = append(fieldList, field)
		}
		sort.Strings(fieldList)

		// Simplify type name (remove package path for readability)
		shortType := typeName
		if idx := strings.LastIndex(typeName, "."); idx >= 0 {
			shortType = typeName[idx+1:]
		}

		parts = append(parts, shortType+"{"+strings.Join(fieldList, ", ")+"}")
	}
	sort.Strings(parts)

	return funcName + ": " + strings.Join(parts, ", ")
}

// hasDirective checks if a comment group contains a specific directive
func hasDirective(comments *ast.CommentGroup, directive string) bool {
	if comments == nil {
		return false
	}

	for _, comment := range comments.List {
		if strings.Contains(comment.Text, directive) {
			return true
		}
	}

	return false
}
