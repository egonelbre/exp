package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var keywords = []string{
		"break", "case", "chan", "const", "continue",
		"default", "defer", "else", "fallthrough", "for", "func", "go",
		"goto", "if", "import", "interface", "map", "package", "range", "return",
		"select", "struct", "switch", "type", "var",
	}

	allfuncs := []string{}

	generators := []Generator{
		LengthSingle{},
		TwoHashShift{},
		TwoHashShiftTable{},
	}
	for _, gen := range generators {
		genname := strings.ToLower(fmt.Sprintf("%+T", gen))
		genname = strings.TrimPrefix(genname, "main.")

		p := &BasicPrinter{}
		p.F("package check\n\n")
		gen.Generate(p, append([]string{}, keywords...))

		emitfile(filepath.Join("check", genname)+".go", p.Bytes())

		allfuncs = append(allfuncs, p.FuncNames()...)
	}

	{
		p := &BasicPrinter{}
		p.F("package main\n\n")
		p.F("import . \"github.com/egonelbre/exp/brutecheck/check\"\n\n")
		for _, fn := range allfuncs {
			p.F("func Bench%s(words []string, repeat int) int {\n", fn)
			p.F("var count int\n")
			p.F("for i := 0; i < repeat; i++ {\n")
			p.F("for _, word := range words {\n")
			p.F("if %s(word) { count++ }\n", fn)
			p.F("}\n")
			p.F("}\n")
			p.F("return count\n")
			p.F("}\n\n")
		}
		emitfile(filepath.Join("bench", "bench.go"), p.Bytes())
	}

	{
		p := &BasicPrinter{}
		p.F("package main\n\n")
		p.F("type Bench struct { Name string; Fn func([]string, int) int }\n\n")
		p.F("var AllBenches = []Bench{\n")
		for _, fn := range allfuncs {
			p.F("{Name:%q, Fn: Bench%s},\n", fn, fn)
		}
		p.F("}")
		emitfile(filepath.Join("bench", "list.go"), p.Bytes())
	}
}

func emitfile(name string, data []byte) {
	formatted, err := format.Source(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to format: %v\n", err)
		fmt.Fprintf(os.Stderr, "-----")
		if len(data) > 1000 {
			data = data[:1000]
		}
		fmt.Fprint(os.Stderr, string(data))
		os.Exit(1)
	}

	err = ioutil.WriteFile(name, formatted, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write: %v\n", err)
		os.Exit(1)
	}
}

type Generator interface {
	Generate(p Printer, keywords []string)
}

type Printer interface {
	FuncName(fnname string)
	FuncNames() []string

	F(format string, args ...interface{})
	Bytes() []byte
}

type BasicPrinter struct {
	Funcs  []string
	Buffer bytes.Buffer
}

func (p *BasicPrinter) FuncName(fnname string) {
	p.Funcs = append(p.Funcs, fnname)
}

func (p *BasicPrinter) FuncNames() []string {
	return p.Funcs
}

func (p *BasicPrinter) F(format string, args ...interface{}) {
	fmt.Fprintf(&p.Buffer, format, args...)
}

func (p *BasicPrinter) Bytes() []byte {
	return p.Buffer.Bytes()
}
