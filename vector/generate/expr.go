package main

import (
	"regexp"
)

var rxVariable = regexp.MustCompile("\\$[a-zA-Z]+")

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
