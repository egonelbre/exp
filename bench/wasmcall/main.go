package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("add", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return Add(args[0].Float(), args[1].Float())
	}))
	js.Global().Set("sub", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return Sub(args[0].Float(), args[1].Float())
	}))
	select{}
}

//export add
func Add(a, b float64) float64 {
	return a + b
}

//export sub
func Sub(a, b float64) float64 {
	return a - b
}