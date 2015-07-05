package main

import (
	"fmt"
	"time"

	"github.com/egonelbre/exp/zerostream/op"
)

func main() {
	x := op.Stream{make([]byte, 1<<30)}
	it := &op.Iterator{x, 0}

	start := time.Now()
	for i := 0; i < 1<<20; i += 1 {
		*it.Translate() = op.Translate{10, 10}
		it.Start()
		*it.MoveTo() = op.MoveTo{5, 5}
		*it.LineTo() = op.LineTo{20, 10}
		*it.LineTo() = op.LineTo{5, 10}
		it.Close()
		*it.Translate() = op.Translate{-10, -10}
	}
	fmt.Println(time.Since(start))
	fmt.Println("Last ", it.Head)

	start = time.Now()
	it = &op.Iterator{x, 0}
	z := int32(0)
	for !it.EOF() {
		switch it.Type() {
		case op.TypeStart:
			it.Start()
		case op.TypeClose:
			it.Close()
		case op.TypeMoveTo:
			z += it.MoveTo().X
		case op.TypeLineTo:
			z += it.LineTo().X
		case op.TypeTranslate:
			z += it.Translate().Dx
		}
	}
	fmt.Println(time.Since(start))
	fmt.Println("Last ", it.Head, z)
}
