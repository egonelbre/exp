package main

import "reflect"

//gistsnip:start:ops
type Op struct {
	Name  string
	Left  reflect.Type
	Right reflect.Type
}

var Ops = map[Op]func(a, b Field) Field{}

func init() {
	tUint := reflect.TypeOf(&Uint{})
	tFloat := reflect.TypeOf(&Float{})

	Ops[Op{"Add", tUint, tUint}] = func(a, b Field) Field { return a.(*Uint).Add(b.(*Uint)) }
	Ops[Op{"Sub", tUint, tUint}] = func(a, b Field) Field { return a.(*Uint).Sub(b.(*Uint)) }

	Ops[Op{"Add", tFloat, tFloat}] = func(a, b Field) Field { return a.(*Float).Add(b.(*Float)) }
	Ops[Op{"Sub", tFloat, tFloat}] = func(a, b Field) Field { return a.(*Float).Sub(b.(*Float)) }
}

func Call(name string, a, b Field) Field {
	if aerr, ok := a.(*Error); ok {
		return aerr
	}
	if berr, ok := b.(*Error); ok {
		return berr
	}

	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	call, found := Ops[Op{"Add", ta, tb}]
	if !found {
		return &Error{"unhandled op"}
	}
	return call(a, b)
}

func Add(a, b Field) Field { return Call("Add", a, b) }
func Sub(a, b Field) Field { return Call("Sub", a, b) }

//gistsnip:end:ops
