package main

import (
	"fmt"
	"time"

	"github.com/egonelbre/exp/zerostream/op"

	"github.com/davecheney/profile"
)

type Rect struct{ X, Y, W, H int32 }

func (r Rect) Draw(out *op.Iterator) {
	out.Start()
	*out.MoveTo() = op.MoveTo{r.X, r.Y}
	*out.LineTo() = op.LineTo{r.X + r.W, r.Y}
	*out.LineTo() = op.LineTo{r.X + r.W, r.Y + r.H}
	*out.LineTo() = op.LineTo{r.X, r.Y + r.H}
	out.Close()
}

func (r Rect) Render(out *op.Iterator) {
	r.Draw(out)
	out.Fill()
	out.Stroke()
}

func main() {
	x := op.Stream{make([]byte, 1<<30)}
	it := &op.Iterator{x, 0}

	defer profile.Start(&profile.Config{
		CPUProfile:  true,
		ProfilePath: ".",
	}).Stop()

	start := time.Now()
	for i := 0; i < 1<<20; i += 1 {
		*it.Translate() = op.Translate{2, 2}
		Rect{10, 10, 20, 20}.Render(it)
	}
	fmt.Println(time.Since(start))
	fmt.Println("Last ", it.Head)

	start = time.Now()
	it = &op.Iterator{x, 0}
	z := int32(0)

RENDER:
	for {
		switch it.Type() {
		case op.EOF:
			break RENDER
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
		case op.TypeFill:
			it.Fill()
		case op.TypeStroke:
			it.Stroke()
		default:
			panic("unhandled type")
		}
	}
	fmt.Println(time.Since(start))
	fmt.Println("Last ", it.Head, z)
}
