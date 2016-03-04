package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

type float float64

const (
	tau   = float(math.Pi * 2)
	pixr  = float(1.0)
	pixr2 = pixr * pixr
	pixa  = pixr2 * tau / 2
)

func (v float) Sqrt() float { return float(math.Sqrt(float64(v))) }
func (v float) Sqr() float  { return v * v }
func (v float) Sin() float  { return float(math.Sin(float64(v))) }
func (v float) Cos() float  { return float(math.Cos(float64(v))) }
func (v float) Acos() float { return float(math.Acos(float64(v))) }
func (v float) Abs() float {
	if v < 0 {
		return -v
	}
	return v
}

type Point struct{ X, Y float }
type Line struct {
	A, B   Point
	Radius float
}

func Dist2(a, b Point) float {
	return (a.Y - b.Y).Sqr() + (a.X - b.X).Sqr()
}

func (line *Line) DistanceTo(p Point) (center float) {
	v, w := line.A, line.B
	l2 := Dist2(v, w)
	t := ((p.X-v.X)*(w.X-v.X) + (p.Y-v.Y)*(w.Y-v.Y)) / l2

	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	return Dist2(p, Point{X: v.X + t*(w.X-v.X), Y: v.Y + t*(w.Y-v.Y)}).Sqrt()
}

func SegmentArea(d float) float {
	//if d >= pixr {
	//	return 0
	//}
	//if d <= -pixr {
	//	return pixa
	//}

	t := 2 * (d / pixr).Acos()
	return (1 / 2) * pixr2 * (t - t.Sin())
}

func (line *Line) Coverage(p Point) float {
	d := line.DistanceTo(p)
	if d > line.Radius+pixr {
		return 0
	}

	min, max := d-line.Radius, d+line.Radius
	if min <= -pixr && pixr <= max {
		return 1
	}

	if pixr <= max {
		fmt.Println(">", -pixr, min, d, max, pixr)
		if min < 0 {
			return (pixa - SegmentArea(-min)) / pixa
		}
		return SegmentArea(min) / pixa
	}
	if min <= -pixr {
		fmt.Println("<", -pixr, min, d, max, pixr)
		return SegmentArea(max) / pixa
	}
	fmt.Println("q", -pixr, min, d, max, pixr)
	return (SegmentArea(min) - SegmentArea(max)) / pixa
}

func CoverageToColor(c float) uint8 {
	if c >= 1 {
		return 0x00
	} else if c <= 0 {
		return 0xFF
	} else {
		return 0xFF - uint8(0xFF*c)
	}
}

func main() {
	a := Line{Point{100, 100}, Point{213, 217}, 10}
	w, h := 1024, 1024

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := Point{float(x), float(y)}
			//c := a.Coverage(p)
			c := 10 - a.DistanceTo(p)
			q := CoverageToColor(c)
			m.SetRGBA(x, y, color.RGBA{q, q, q, 0xFF})
		}
	}

	file, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, m); err != nil {
		log.Fatal(err)
	}
}
