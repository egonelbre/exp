package main

import (
	"fmt"
	"strconv"
)

type Config struct {
	Package string

	Unroll  int
	Type    Type
	Counter bool
}

func (c Config) GetConfig() Config { return c }

type Context interface {
	GetConfig() Config

	In(output string) File
}

type File interface {
	GetConfig() Config

	Func(
		signature string,
		template Template,
	)
}

type Type struct {
	Name string
	Size int64
}

func Range(name string, start, count any) It {
	return It{
		Kind:  ItRange,
		Name:  name,
		Start: ExprFrom(start),
		Inc:   Expr{Const: 1},
		Count: ExprFrom(count),
	}
}

func Slice(name string, start, inc any) It {
	return It{
		Kind:  ItSlice,
		Name:  name,
		Start: ExprFrom(start),
		Inc:   ExprFrom(inc),
		Count: Expr{Derived: true},
	}
}

type Template interface {
	isTemplate()
}

type Iterate struct {
	Ranges []It
	For    string
}

func (Iterate) isTemplate() {}

type It struct {
	Kind  ItKind
	Name  string
	Start Expr
	Inc   Expr
	Count Expr
}

type ItKind byte

const (
	ItRange = 1
	ItSlice = 2
)

type Expr struct {
	Const   int
	Expr    string
	Derived bool
}

func Var(v string) Expr {
	return Expr{Expr: v}
}

func Const(v int) Expr {
	return Expr{Const: v}
}

func ExprFrom(v any) Expr {
	switch v := v.(type) {
	case int:
		return Expr{Const: v}
	case string:
		return Expr{Expr: v}
	default:
		panic(fmt.Sprintf("unhandled %T", v))
	}
}

func (num Expr) String() string {
	switch {
	case num.Derived:
		panic("cannot automatically derive")
	case num.Expr != "":
		return num.Expr
	default:
		return strconv.Itoa(num.Const)
	}
}
