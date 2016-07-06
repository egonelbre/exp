package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"time"
)

func main() {
	save("shade.png", Shade)
}

const (
	ZOOM         = 0.8
	VOLUME_STEPS = 20
	ITERATIONS   = 17

	FORMULA_PARAM   = 0.53
	BRIGHTNESS      = 0.0015
	DISTANCE_FADING = 0.730
	STEP_SIZE       = 0.1
)

func Shade(p, uv, res V2) Color {
	uv.Y *= res.Y / res.X

	dir := V3{uv.X * ZOOM, uv.Y * ZOOM, 1}
	from := V3{0.98, 0.5122, 0.53}

	s := float(0.1)
	fade := float(1)

	v := V3{}
	for r := 0; r < VOLUME_STEPS; r++ {
		p := from.Add(dir.Scale(s * 0.5))
		var pa, a float
		for i := 0; i < ITERATIONS; i++ {
			p = p.Abs().Scale(1.0 / p.Len2()).Offset(-FORMULA_PARAM)
			a += abs(p.Len() - pa)
		}
		a *= a * a
		v = v.Offset(fade)
		v = v.Add(V3{s, s * s, s * s * s * s}.Scale(a * fade * BRIGHTNESS))

		fade *= DISTANCE_FADING
		s += STEP_SIZE
	}

	v = v.Scale(0.01)
	return C3(v.X, v.Y, v.Z)
}

type Shader func(p, uv V2, res V2) Color

func Render(m *image.RGBA, shader Shader) {
	size := m.Bounds().Size()
	res := V2{float(size.X), float(size.Y)}
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			p := V2{float(x), float(y)}
			uv := V2{float(x) / res.X, float(y) / res.Y}
			color := shader(p, uv, res)
			m.SetRGBA(x, y, color.RGBA())
		}
	}
}

func save(name string, shader Shader) {
	fmt.Println("RENDER ", name)
	start := time.Now()
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	Render(m, shader)
	stop := time.Now()
	fmt.Println("DONE ", stop.Sub(start))

	f, err := os.Create(name)
	check(err)
	check(png.Encode(f, m))
	check(f.Close())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// vector stuff

type V2 struct{ X, Y float }
type V3 struct{ X, Y, Z float }

func (v V2) Abs() V2 { return V2{abs(v.X), abs(v.Y)} }
func (v V3) Abs() V3 { return V3{abs(v.X), abs(v.Y), abs(v.Z)} }

func (v V2) Len2() float { return v.X*v.X + v.Y*v.Y }
func (v V3) Len2() float { return v.X*v.X + v.Y*v.Y + v.Z*v.Z }

func (v V2) Len() float { return sqrt(v.Len2()) }
func (v V3) Len() float { return sqrt(v.Len2()) }

func (v V2) LenV() V2 {
	d := v.Len()
	return V2{d, d}
}
func (v V3) LenV() V3 {
	d := v.Len()
	return V3{d, d, d}
}

func (v V2) Scale(s float) V2 { return V2{v.X * s, v.Y * s} }
func (v V3) Scale(s float) V3 { return V3{v.X * s, v.Y * s, v.Z * s} }

func (a V2) Add(b V2) V2 { return V2{a.X + b.X, a.Y + b.Y} }
func (a V3) Add(b V3) V3 { return V3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a V2) Offset(s float) V2 { return V2{a.X + s, a.Y + s} }
func (a V3) Offset(s float) V3 { return V3{a.X + s, a.Y + s, a.Z + s} }

// general math

const (
	Pi  = math.Pi
	Tau = 2 * math.Pi
)

type float float64

func abs(v float) float {
	if v < 0 {
		return -v
	}
	return v
}

func cos(v float) float  { return float(math.Cos(float64(v))) }
func sin(v float) float  { return float(math.Sin(float64(v))) }
func sqrt(v float) float { return float(math.Sqrt(float64(v))) }

func lerp(a, b, p float) float { return (1.0-p)*a + p*b }
func u8(v float) uint8 {
	x := int(v * 256)
	if x >= 256 {
		return 255
	} else if x < 0 {
		return 0
	}
	return uint8(x)
}

func mod1(v float) float {
	vi := float(math.Trunc(float64(v)))
	if vi < 0 {
		return v - vi + 1
	}
	return v - vi
}

func clamp1(v float) float {
	if v > 1 {
		return 1
	} else if v < 0 {
		return 0
	}
	return v
}

// Color utilities
type Color struct{ R, G, B, A float }

func (c Color) RGBA() color.RGBA {
	return color.RGBA{u8(c.R), u8(c.G), u8(c.B), u8(c.A)}
}

func G(gray float) Color        { return Color{gray, gray, gray, 1} }
func C3(R, G, B float) Color    { return Color{R, G, B, 1} }
func C4(R, G, B, A float) Color { return Color{R, G, B, A} }

func hue(p, q, h float) float {
	if h < 0.0 {
		h += 1.0
	} else if h > 1.0 {
		h -= 1.0
	}
	if h < 1.0/6.0 {
		return p + (q-p)*6.0*h
	} else if h < 1.0/2.0 {
		return q
	} else if h < 2.0/3.0 {
		return p + (q-p)*6.0*(2.0/3.0-h)
	}
	return p
}

func HSL(H, S, L float) Color {
	S, L = clamp1(S), clamp1(L)
	if S <= 0 {
		return C3(L, L, L)
	}
	H = mod1(H)

	var q float
	if L < 0.5 {
		q = L * (1 + S)
	} else {
		q = L + S - S*L
	}
	p := 2*L - q
	return C3(hue(p, q, H+1.0/3.0), hue(p, q, H), hue(p, q, H-1.0/3.0))
}

func HSLA(H, S, L, A float) Color {
	c := HSL(H, S, L)
	c.A = A
	return c
}
