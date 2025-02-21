package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"maps"
	"slices"

	"golang.org/x/tools/go/packages"
)

type Category struct {
	Size        int64
	HasPointers bool
}

func main() {
	top := flag.Int("top", 10, "only print top entries")

	flag.Parse()

	config := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.LoadImports | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedTypesSizes,
	}

	roots, err := packages.Load(config, flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	summary := map[Category]int64{}
	summaryByType := map[types.Type]int64{}

	for _, pkg := range All(roots) {
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				call, ok := n.(*ast.CallExpr)
				if !ok {
					return true
				}

				if !isAppendName(call.Fun) {
					return true
				}
				if call.Args[0] == nil {
					return true
				}
				argType := pkg.TypesInfo.TypeOf(call.Args[0])
				if argType == nil {
					return true
				}

				argType = argType.Underlying()

				sliceType, ok := argType.(*types.Slice)
				if !ok {
					ast.Fprint(os.Stderr, pkg.Fset, call, ast.NotNilFilter)
					fmt.Fprintf(os.Stderr, "%#v as argument\n", argType)
					panic("invalid argument to append")
				}

				elemType := sliceType.Elem()
				if _, isGeneric := elemType.(*types.TypeParam); isGeneric {
					return true
				}

				size := pkg.TypesSizes.Sizeof(elemType)
				cat := Category{
					Size:        size,
					HasPointers: HasPointers(elemType),
				}
				summary[cat]++
				summaryByType[elemType]++

				return true
			})
		}
	}

	{
		fmt.Println("TOP BY CATEGORY")
		results := KeyValues(summary)
		slices.SortFunc(results, func(a, b KV[Category, int64]) bool {
			return a.Value > b.Value
		})
		for _, r := range Top(results, *top) {
			ptr := ""
			if r.Key.HasPointers {
				ptr = "ptr"
			}
			fmt.Printf("  %6v %4v %v bytes\n", r.Value, ptr, r.Key.Size)
		}
	}
	{
		fmt.Println("TOP BY TYPE")
		var results []KV[types.Type, int64]
		for k, v := range summaryByType {
			results = append(results, KV[types.Type, int64]{Key: k, Value: v})
		}
		slices.SortFunc(results, func(a, b KV[types.Type, int64]) bool {
			return a.Value > b.Value
		})
		for _, r := range Top(results, *top) {
			ptr := ""
			if HasPointers(r.Key) {
				ptr = "ptr"
			}
			fmt.Printf("  %6v %4v %v\n", r.Value, ptr, r.Key)
		}
	}
}

func All(roots []*packages.Package) []*packages.Package {
	all := map[*packages.Package]struct{}{}

	var walk func(*packages.Package)
	walk = func(p *packages.Package) {
		if _, seen := all[p]; seen {
			return
		}
		all[p] = struct{}{}

		for _, c := range p.Imports {
			walk(c)
		}
	}
	for _, root := range roots {
		walk(root)
	}

	pkgs := maps.Keys(all)
	slices.SortFunc(pkgs, func(a, b *packages.Package) bool {
		return a.ID < b.ID
	})
	return pkgs
}

func Top[T any](xs []T, n int) []T {
	if len(xs) <= n {
		return xs
	}
	return xs[:n]
}

type KV[K, V any] struct {
	Key   K
	Value V
}

func KeyValues[K comparable, V any](m map[K]V) []KV[K, V] {
	var kvs []KV[K, V]
	for k, v := range m {
		kvs = append(kvs, KV[K, V]{Key: k, Value: v})
	}
	return kvs
}

func isAppendName(x ast.Expr) bool {
	switch x := x.(type) {
	case *ast.SelectorExpr:
		return false
	case *ast.Ident:
		return x.Name == "append"
	default:
		return false
	}
}

func qualifiedName(x ast.Expr) string {
	switch x := x.(type) {
	case *ast.SelectorExpr:
		pkg, ok := x.X.(*ast.Ident)
		if !ok {
			return ""
		}
		return pkg.Name + "." + x.Sel.Name
	case *ast.Ident:
		return x.Name
	default:
		return ""
	}
}

var hasPointers = map[types.Type]bool{}

func HasPointers(T types.Type) (result bool) {
	TU := T.Underlying()
	if cached, ok := hasPointers[TU]; ok {
		return cached
	}
	hasPointers[TU] = true
	defer func() { hasPointers[TU] = result }()

	switch t := TU.(type) {
	case *types.Basic:
		switch t.Kind() {
		case types.String, types.UnsafePointer:
			return true
		}
		return false
	case *types.Chan, *types.Map, *types.Pointer, *types.Signature, *types.Slice:
		return true
	case *types.Interface:
		return true
	case *types.Array:
		return t.Len() > 0 && HasPointers(t.Elem())
	case *types.Struct:
		nf := t.NumFields()
		for i := 0; i < nf; i++ {
			if HasPointers(t.Field(i).Type()) {
				return true
			}
		}
		return false
	}

	panic("impossible")
}
