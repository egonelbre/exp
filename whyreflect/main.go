package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	deps := Parse(os.Stdin)

	all := New(deps...)
	reflect := Filter(all, func(b *Node) bool {
		return b.Flags&ReflectMethod != 0
	})
	reach := Reach(all, reflect)

	// sort output
	nodes := maps.Keys(reach)
	slices.SortFunc(nodes, func(a, b *Node) bool { return a.Sym.Name < b.Sym.Name })
	for _, n := range nodes {
		slices.SortFunc(n.Needs, func(a, b *Node) bool { return a.Sym.Name < b.Sym.Name })
	}

	fmt.Print("digraph G {\n")
	defer fmt.Print("}\n")

	fmt.Print("    node [fontsize=10 shape=rectangle target=\"_graphviz\"];\n")
	fmt.Print("    edge [tailport=e];\n")
	fmt.Print("    compound=true;\n")
	fmt.Print("    rankdir=LR;\n")
	fmt.Print("    newrank=true;\n")
	fmt.Print("    ranksep=\"1.5\";\n")
	fmt.Print("    quantum=\"0.5\";\n")

	fmt.Print("\n")
	for n := range reach {
		if n.Flags&ReflectMethod != 0 {
			fmt.Printf("    %v [label=%q, style=filled, fillcolor=lightgray]\n", nodeID(n), n.Name)
		} else {
			fmt.Printf("    %v [label=%q]\n", nodeID(n), n.Name)
		}
	}
	for n := range reach {
		for _, x := range n.Needs {
			if _, ok := reach[x]; !ok {
				continue
			}
			fmt.Printf("    %v -> %v\n", nodeID(n), nodeID(x))
		}
	}
}

func nodeID(node *Node) string {
	return strconv.Quote(node.Name)
}

type Dep struct {
	From Sym
	To   Sym
}

type Sym struct {
	Name  string
	Flags Flag
}

type Flag byte

const (
	UsedInIface   = 1
	ReflectMethod = 2
)

func Parse(r io.Reader) (deps []Dep) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		from, to, ok := strings.Cut(line, " -> ")
		if !ok {
			continue
		}

		deps = append(deps, Dep{
			From: ParseSym(from),
			To:   ParseSym(to),
		})
	}
	return deps
}

func ParseSym(t string) (sym Sym) {
	old := ""
	for old != t {
		old = t
		if strings.HasSuffix(t, " <UsedInIface>") {
			t = strings.TrimSuffix(t, " <UsedInIface>")
			sym.Flags |= UsedInIface
		}
		if strings.HasSuffix(t, " <ReflectMethod>") {
			t = strings.TrimSuffix(t, " <ReflectMethod>")
			sym.Flags |= ReflectMethod
		}
	}
	sym.Name = t
	return sym
}

var token struct{}

type Node struct {
	Sym
	Needs []*Node
}

type Graph map[*Node]struct{}

func New(deps ...Dep) Graph {
	nodes := make(map[Sym]*Node)
	set := make(Graph)

	ensure := func(sym Sym) *Node {
		node, ok := nodes[sym]
		if !ok {
			node = &Node{Sym: sym}
			nodes[sym] = node
			set[node] = token
		}
		return node
	}

	for _, dep := range deps {
		src := ensure(dep.From)
		dst := ensure(dep.To)

		src.Needs = append(src.Needs, dst)
	}

	return set
}

func Filter(a Graph, fn func(b *Node) bool) Graph {
	result := New()
	for n := range a {
		if fn(n) {
			result[n] = token
		}
	}
	return result
}

// Reach returns packages in a that terminate in b.
func Reach(a, b Graph) Graph {
	result := maps.Clone(b)
	candidates := maps.Clone(a)

	for x := range result {
		delete(candidates, x)
	}

	for modified := true; modified; {
		modified = false
		for candidate := range candidates {
			for _, dep := range candidate.Needs {
				if _, ok := result[dep]; ok {
					result[candidate] = token
					delete(candidates, candidate)
					modified = true
				}
			}
		}
	}

	return result
}
