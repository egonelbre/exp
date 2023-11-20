package main

import (
	"bytes"
	"fmt"
	"go/format"
	"regexp"
	"strings"
)

type Naive struct {
	Config

	Files []*NaiveFile
}

func NewNaive(config Config) *Naive {
	return &Naive{Config: config}
}

type NaiveFile struct {
	Config

	Path    string
	Content bytes.Buffer
}

func (ctx *Naive) In(path string) File {
	for _, file := range ctx.Files {
		if file.Path == path {
			return file
		}
	}

	file := &NaiveFile{Path: path, Config: ctx.Config}
	ctx.Files = append(ctx.Files, file)

	file.emitHeader()

	return file
}

func (ctx *NaiveFile) Formatted() ([]byte, error) {
	formatted, err := format.Source(ctx.Content.Bytes())
	if err != nil {
		return ctx.Content.Bytes(), err
	}
	return formatted, nil
}

func (ctx *NaiveFile) emitHeader() {
	ctx.Printf("package %s\n", ctx.Config.Package)
}

func (ctx *NaiveFile) Func(signature string, template Template) {
	switch t := template.(type) {
	case Iterate:
		ctx.Iterate(signature, t)
	default:
		panic(fmt.Sprintf("unhandled %T", t))
	}
}

func (ctx *NaiveFile) Iterate(signature string, body Iterate) {
	pf := ctx.Printf

	pf("\n")
	pf("func %s {\n", ctx.specializeSignature(signature))
	defer pf("}\n")

	// determine the primary iterator
	itCandidates := []It{}
	for _, it := range body.Ranges {
		if it.Count.Expr != "" {
			itCandidates = append(itCandidates, it)
		}
	}

	// maybe generate a primary iterator based on the slices we have
	firstRangeIsPrimary := false
	if len(itCandidates) == 0 {
		firstRangeIsPrimary = true
		first := body.Ranges[0]

		ensure(first.Inc.Const == 1 && first.Inc.Expr == "")
		ensure(first.Start.Const == 0 && first.Start.Expr == "")

		// generate a range based on the first item in the slice
		pf("n := len(%s)\n", first.Name)
		r := Range("i", 0, "n")
		body.Ranges = append(body.Ranges, r)
		itCandidates = []It{r}
	}
	ensure(len(itCandidates) == 1)

	// generate boundary checks for the iteration
	prim := itCandidates[0]
	for i, it := range body.Ranges {
		if i == 0 && firstRangeIsPrimary {
			// skip the one we used to calculate `n`
			continue
		}
		if !it.Count.Derived {
			continue
		}

		size := Var("len(" + it.Name + ")")
		// TODO: handle multiplication overflow
		pf("if %s < int(%s) { panic(\"%s is too small\") }\n", sub(size, it.Start), mul(prim.Count, it.Inc), it.Name)
	}

	if ctx.Unroll <= 1 {
		if ctx.Counter {
			for i, it := range body.Ranges {
				if i == 0 && firstRangeIsPrimary || it == prim {
					continue
				}

				pf("i%s := %s\n", it.Name, it.Start)
			}
		}

		// TODO: simplify increment
		pf("for %s := %s ; %s < %s; %s {\n", prim.Name, prim.Start, prim.Name, prim.Count, increment(prim.Name, prim.Inc))
		if ctx.Counter {
			pf("	%s\n", ctx.specializeTemplateCounterAccess(body.For, prim, 0, body.Ranges))
		} else {
			pf("	%s\n", ctx.specializeTemplateDirectAccess(body.For, prim, 0, body.Ranges))
		}
		if ctx.Counter {
			for i, it := range body.Ranges {
				if i == 0 && firstRangeIsPrimary || it == prim {
					continue
				}
				pf("	%s\n", increment("i"+it.Name, it.Inc))
			}
		}
		pf("}\n")
	} else {
		pf("%s := %s\n", prim.Name, prim.Start)
		pf("%s_unroll := %s - %s %% %v\n", prim.Count, prim.Count, prim.Count, ctx.Unroll)

		if ctx.Counter {
			for i, it := range body.Ranges {
				if i == 0 && firstRangeIsPrimary || it == prim {
					continue
				}

				pf("i%s := %s\n", it.Name, it.Start)
			}
		}

		pf("for ; %s < %s_unroll; %s {\n", prim.Name, prim.Count, increment(prim.Name, mul(Const(ctx.Unroll), prim.Inc)))
		for i := 0; i < ctx.Unroll; i++ {
			if ctx.Counter {
				pf("	%s\n", ctx.specializeTemplateCounterAccess(body.For, prim, i, body.Ranges))
			} else {
				pf("	%s\n", ctx.specializeTemplateDirectAccess(body.For, prim, i, body.Ranges))
			}
		}
		if ctx.Counter {
			for i, it := range body.Ranges {
				if i == 0 && firstRangeIsPrimary || it == prim {
					continue
				}
				pf("	%s\n", increment("i"+it.Name, mul(it.Inc, Const(ctx.Unroll))))
			}
		}
		pf("}\n")

		pf("for ; %s < %s; %s {\n", prim.Name, prim.Count, increment(prim.Name, prim.Inc))
		if ctx.Counter {
			pf("	%s\n", ctx.specializeTemplateCounterAccess(body.For, prim, 0, body.Ranges))
			for i, it := range body.Ranges {
				if i == 0 && firstRangeIsPrimary || it == prim {
					continue
				}
				pf("	%s\n", increment("i"+it.Name, it.Inc))
			}
		} else {
			pf("	%s\n", ctx.specializeTemplateDirectAccess(body.For, prim, 0, body.Ranges))
		}
		pf("}\n")
	}
}

func (ctx *NaiveFile) specializeSignature(signature string) string {
	return strings.ReplaceAll(signature, "$Type", ctx.Config.Type.Name)
}

var rxVariable = regexp.MustCompile("\\$[a-zA-Z]+")

func (ctx *NaiveFile) specializeTemplateDirectAccess(body string, prim It, primOffset int, ranges []It) string {
	return rxVariable.ReplaceAllStringFunc(body, func(ref string) string {
		ref = ref[1:]

		for _, r := range ranges {
			if r.Name == ref {
				at := ref + "[" + add(r.Start,
					mul(
						add(Var(prim.Name), Const(primOffset)),
						r.Inc,
					),
				).String() + "]"
				return at
			}
		}

		panic("did not find " + ref)
	})
}

func (ctx *NaiveFile) specializeTemplateCounterAccess(body string, prim It, primOffset int, ranges []It) string {
	return rxVariable.ReplaceAllStringFunc(body, func(ref string) string {
		ref = ref[1:]

		for _, it := range ranges {
			if it.Name == ref {
				at := it.Name + "[i" + it.Name
				if primOffset > 0 {
					at += "+" + mul(Const(primOffset), it.Inc).String()
				}
				at += "]"
				return at
			}
		}

		panic("did not find " + ref)
	})
}

func (ctx *NaiveFile) Printf(format string, args ...any) {
	fmt.Fprintf(&ctx.Content, format, args...)
}

func ensure(v bool) {
	if !v {
		panic("unexpected")
	}
}

// TODO: handle operation order

func mul(a, b Expr) Expr {
	as := a.String()
	bs := b.String()
	if as == "1" {
		return Expr{Expr: bs}
	}
	if bs == "1" {
		return Expr{Expr: as}
	}
	return Expr{Expr: as + " * " + bs}
}

func add(a, b Expr) Expr {
	as := a.String()
	bs := b.String()
	if as == "0" {
		return Expr{Expr: bs}
	}
	if bs == "0" {
		return Expr{Expr: as}
	}
	return Expr{Expr: as + " + " + bs}
}

func sub(a, b Expr) Expr {
	as := a.String()
	bs := b.String()
	if as == "0" {
		return Expr{Expr: "-" + bs}
	}
	if bs == "0" {
		return Expr{Expr: as}
	}
	return Expr{Expr: as + " - " + bs}
}

func increment(varname string, by Expr) string {
	bys := by.String()
	switch bys {
	case "1":
		return varname + "++"
	}
	return varname + " += " + bys
}
