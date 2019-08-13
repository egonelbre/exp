package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"io"
	"os"
	"strconv"
	"strings"
)

var template = `
func () {
	if _, νok := νidx[νkey]; νok {
	    νidx[νkey] = append(νidx[νkey], νvalue)
	} else {
		νidx[νkey] = []τvalue{νvalue}
	}
}
`

func main() {
	t, err := parser.ParseExpr(template)
	fmt.Println(err)
	root := t.(*ast.FuncLit).Body.List[0]

	var b bytes.Buffer
	fmt.Fprintf(&b, "package x\nfunc match(p interface{}, varlist, typelist map[string][]*ast.Ident) bool{\n")
	WriteMatch(&b, "p", root, false)
	fmt.Fprintf(&b, "}\n")

	formatted, err := format.Source(b.Bytes())
	fmt.Println(err)
	if err != nil {
		fmt.Println("ERRORED:", b.String())
	}
	fmt.Println(string(formatted))

	ast.Print(nil, root)
}

func WriteMatch(w io.Writer, src string, e ast.Node, skip bool) {
	dst := sanitize(src)

	if !skip {
		fmt.Fprintf(w, "\t%s, ok := %s.(%T);\n\tif !ok { return false }\n", dst, src, e)
	} else {
		dst = src
	}

	switch e := e.(type) {
	case *ast.AssignStmt:
		fmt.Fprintf(w, "\tif len(%s.Lhs) != %d || len(%s.Rhs) != %d { return false }\n", dst, len(e.Lhs), dst, len(e.Rhs))
		for idx := range e.Lhs {
			WriteMatch(w, dst+".Lhs["+strconv.Itoa(idx)+"]", e.Lhs[idx], false)
		}
		for idx := range e.Rhs {
			WriteMatch(w, dst+".Rhs["+strconv.Itoa(idx)+"]", e.Rhs[idx], false)
		}

	case *ast.CompositeLit:
		WriteMatch(w, dst+".Type", e.Type, false)
		fmt.Fprintf(w, "\tif len(%s.Elts) != %d { return false }\n", dst, len(e.Elts))
		for idx := range e.Elts {
			WriteMatch(w, dst+".Elts["+strconv.Itoa(idx)+"]", e.Elts[idx], false)
		}
	case *ast.ArrayType:
		WriteMatch(w, dst+".Elt", e.Elt, false)
	case *ast.BasicLit:
		// TODO: use proper constants for .Kind
		fmt.Fprintf(w, "\tif %s.Kind != %d || %s.Value != %q { return false }\n", dst, e.Kind, dst, e.Value)
	case *ast.Ident:
		if strings.HasPrefix(e.Name, "ν") {
			name := strings.TrimPrefix(e.Name, "ν")
			fmt.Fprintf(w, "\tvarlist[%q] = append(varlist[%q], %s)\n", name, name, dst)
		} else if strings.HasPrefix(e.Name, "τ") {
			name := strings.TrimPrefix(e.Name, "τ")
			fmt.Fprintf(w, "\ttypelist[%q] = append(typelist[%q], %s)\n", name, name, dst)
		} else {
			fmt.Fprintf(w, "\tif %q != %s.Name { return false }\n", e.Name, dst)
		}
	case *ast.ExprStmt:
		WriteMatch(w, dst+".X", e.X, false)
	case *ast.CallExpr:
		WriteMatch(w, dst+".Fun", e.Fun, false)
		fmt.Fprintf(w, "\tif len(%s.Args) != %d { return false }\n", dst, len(e.Args))
		for idx := range e.Args {
			WriteMatch(w, dst+".Args["+strconv.Itoa(idx)+"]", e.Args[idx], false)
		}
	case *ast.IndexExpr:
		WriteMatch(w, dst+".X", e.X, false)
		WriteMatch(w, dst+".Index", e.Index, false)
	case *ast.BinaryExpr:
		// TODO: use proper constants for .Op
		fmt.Fprintf(w, "\tif %s.Op != %d { return false }\n", dst, e.Op)
		WriteMatch(w, dst+".X", e.X, false)
		WriteMatch(w, dst+".Y", e.Y, false)
	case *ast.ReturnStmt:
		fmt.Fprintf(w, "\tif len(%s.Results) != %d { return false }\n", dst, len(e.Results))
		for idx := range e.Results {
			WriteMatch(w, dst+".Results["+strconv.Itoa(idx)+"]", e.Results[idx], false)
		}
	case *ast.IfStmt:
		WriteMatch(w, dst+".Init", e.Init, false)
		WriteMatch(w, dst+".Cond", e.Cond, false)
		WriteMatch(w, dst+".Body", e.Body, true)
		WriteMatch(w, dst+".Else", e.Else, false)
	case *ast.BlockStmt:
		fmt.Fprintf(w, "\tif len(%s.List) != %d { return false }\n", dst, len(e.List))
		for idx := range e.List {
			WriteMatch(w, dst+".List["+strconv.Itoa(idx)+"]", e.List[idx], false)
		}
	default:
		fmt.Fprintf(os.Stderr, "unhandled %T\n", e)
	}
}

func sanitize(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '.', '[', ']':
			return '_'
		default:
			return r
		}
	}, s)
}
