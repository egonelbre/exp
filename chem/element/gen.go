// +build ignore

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Element struct {
	AtomicNumber int
	Symbol       string
	Name         string
}

func main() {
	file, err := os.Open("elements.csv")
	if err != nil {
		log.Fatal(err)
	}

	data := csv.NewReader(file)
	records, err := data.ReadAll()
	if err != nil || len(records) == 0 {
		log.Fatal(err)
	}

	var elements []Element
	for _, rec := range records[1:] {
		nr, err := strconv.Atoi(rec[0])
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, Element{
			AtomicNumber: nr,
			Symbol:       strings.TrimSpace(rec[1]),
			Name:         strings.TrimSpace(rec[2]),
		})
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Symbol < elements[j].Symbol
	})

	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package element\n")
	fmt.Fprintf(&buf, "type Element byte\n")

	fmt.Fprintf(&buf, "const (\n")
	fmt.Fprintf(&buf, "\tInvalid = Element(0)\n")
	for i, elem := range elements {
		fmt.Fprintf(&buf, "\t%v = Element(%v)\n", elem.Symbol, i+1)
	}
	fmt.Fprintf(&buf, ")\n")

	fmt.Fprintf(&buf, "var AtomSymbol = [...]string{\"<INVALID>\",\n")
	for i, elem := range elements {
		fmt.Fprintf(&buf, "%q, ", elem.Symbol)
		if (i+1)%10 == 0 {
			fmt.Fprintf(&buf, "\n")
		}
	}
	fmt.Fprintf(&buf, "\n}\n")

	fmt.Fprintf(&buf, "func (e Element) String() string {\n")
	fmt.Fprintf(&buf, "\treturn AtomSymbol[int(e)]\n")
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "func (m *Matcher) acceptElement(p, t byte) {\n")
	fmt.Fprintf(&buf, "p0, p1, p2 := m.src[p], m.src[p+1], m.src[p+2]\n") // "

	root := Trie{}
	for _, elem := range elements {
		root.Add(elem.Symbol, strings.ToLower(elem.Symbol))
	}

	var emitLayer func(t *Trie, n int)
	emitLayer = func(t *Trie, n int) {
		if t.Accept != "" {
			defer fmt.Fprintf(&buf, "m.accepted(%v, p+%v, t)\n", t.Accept, n)
		}

		if len(t.Child) == 0 {
			return
		}

		if len(t.Child) == 1 {
			fmt.Fprintf(&buf, "if p%v == %q {\n", n, t.Child[0].Rune)
			defer fmt.Fprintf(&buf, "}\n")
			emitLayer(t.Child[0], n+1)
			return
		}

		fmt.Fprintf(&buf, "switch p%v {\n", n)
		defer fmt.Fprintf(&buf, "}\n")

		for _, child := range t.Child {
			fmt.Fprintf(&buf, "case %q:\n", child.Rune)
			emitLayer(child, n+1)
		}
	}
	emitLayer(&root, 0)
	fmt.Fprintf(&buf, "}\n\n")

	//ioutil.WriteFile("element.raw.go", buf.Bytes(), 0755)

	srcdata, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("element.go", srcdata, 0755)
}

type Trie struct {
	Rune   rune
	Accept string
	Child  []*Trie
}

func (t *Trie) Add(elem string, s string) {
	if s == "" {
		t.Accept = elem
		return
	}

	for _, n := range t.Child {
		if n.Rune == rune(s[0]) {
			n.Add(elem, s[1:])
			return
		}
	}

	n := &Trie{}
	n.Rune = rune(s[0])
	n.Add(elem, s[1:])
	t.Child = append(t.Child, n)
}
