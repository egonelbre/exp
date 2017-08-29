package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Flowers!",
		Bounds:    pixel.R(0, 0, 480, 520),
		Resizable: false,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	canvas := imdraw.New(nil)
	canvas.Color = color.RGBA{0x00, 0x00, 0x00, 0x80}
	{ // internal line
		y := 64.0
		for line := 1.0; line < 64; line *= 2 {
			canvas.Push(
				pixel.V(50, y),
				pixel.V(100, y-16),
				pixel.V(150, y+16),
				pixel.V(200, y-16),
				pixel.V(100, y-64),
			)
			canvas.Line(line)
			y += 80
		}
	}

	{ // better line
		y := 64.0
		for line := 1.0; line < 64; line *= 2 {
			drawLine(
				canvas,
				line,
				pixel.V(240+50, y),
				pixel.V(240+100, y-16),
				pixel.V(240+150, y+16),
				pixel.V(240+200, y-16),
				pixel.V(240+100, y-64),
			)
			y += 80
		}
	}

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
		canvas.Draw(win)
		win.Update()
	}
}

func drawLine(m *imdraw.IMDraw, width float64, path ...pixel.Vec) {
	if len(path) < 2 {
		return
	}

	radius := width / 2

	// draw each segment, where
	//
	// x1--^------a1-------^----------b1
	//     | xn   |        | abn      |
	// -----------a - - - - - - - - - b
	//            |                   |
	// x2---------a2------------------b2
	//             drawing this segment
	//            |-------------------|
	//            ^__ this will contain a bend
	// x1, x2, xn are the previous segments end
	// corners and normal

	a := path[0]
	var x1, x2, xn pixel.Vec
	for i, b := range path {
		if i > 0 && a.Sub(b).Len() < 1 {
			continue
		}
		// calculate
		abn := ScaleTo(SegmentNormal(a, b), radius)
		// segment-corners
		a1, a2 := a.Add(abn), a.Sub(abn)
		b1, b2 := b.Add(abn), b.Sub(abn)

		// fill the gap between bends
		if i > 0 && radius > 1.5 {
			// see which direction is the gap in
			d := Rotate(xn).Dot(abn)
			if d < 0 {
				m.Push(x1, a1, a)
			} else {
				m.Push(x2, a2, a)
			}
			m.Polygon(0)
		}

		// draw the segment
		m.Push(a1, b1, b2, a2)
		m.Polygon(0)

		// update points
		a = b
		x1, x2, xn = b1, b2, abn
	}
}

func drawClosedLine(m *imdraw.IMDraw, width float64, path ...pixel.Vec) {
	if len(path) < 2 {
		return
	}

	radius := width / 2

	// draw each segment, where
	//
	// x1--^------a1-------^----------b1
	//     | xn   |        | abn      |
	// -----------a - - - - - - - - - b
	//            |                   |
	// x2---------a2------------------b2
	//             drawing this segment
	//            |-------------------|
	//            ^__ this will contain a bend
	// x1, x2, xn are the previous segments end
	// corners and normal
	a := path[len(path)-1]
	xn := ScaleTo(SegmentNormal(path[len(path)-2], a), radius)
	x1, x2 := a.Add(xn), a.Sub(xn)

	for i, b := range path {
		if i > 0 && a.Sub(b).Len() < 1 {
			continue
		}
		// calculate
		abn := ScaleTo(SegmentNormal(a, b), radius)
		// segment-corners
		a1, a2 := a.Add(abn), a.Sub(abn)
		b1, b2 := b.Add(abn), b.Sub(abn)

		// fill the gap between bends
		if radius > 1.5 {
			// see which direction is the gap in
			d := Rotate(xn).Dot(abn)
			if d < 0 {
				m.Push(x1, a1, a)
			} else {
				m.Push(x2, a2, a)
			}
			m.Polygon(0)
		}

		// draw the segment
		m.Push(a1, b1, b2, a2)
		m.Polygon(0)

		// update points
		a = b
		x1, x2, xn = b1, b2, abn
	}
}

const TAU = 2 * math.Pi

func SegmentNormal(a, b pixel.Vec) pixel.Vec {
	return Rotate(b.Sub(a))
}

func Rotate(a pixel.Vec) pixel.Vec {
	return pixel.V(-a.Y, a.X)
}

func ScaleTo(a pixel.Vec, r float64) pixel.Vec {
	x := a.Len()
	return a.Scaled(r / x)
}
