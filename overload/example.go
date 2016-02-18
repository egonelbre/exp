package main

import (
	"fmt"

	g "github.com/egonelbre/exp/overload/template"
)

func Run() {
	var a, b, c, d g.Vector
	fmt.Println(a + b + c + d)
}
