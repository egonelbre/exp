package main

import (
	"fmt"

	"github.com/egonelbre/exp/inline/g"
)

func main() {
	a := g.Vec2{X: 10, Y: 10}
	b := g.Vec2{X: 20, Y: 20}
	// c := a.Add(b).Scale(6) // will be inline into
	c := g.Vec2{X: a.X + b.X, Y: a.Y + b.Y}.Scale(6)
	fmt.Println(c)
}
