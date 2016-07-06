package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	mr "math/rand"
	"os"

	m "github.com/go-gl/mathgl/mgl64"
)

var (
	iterations         int     = 17
	magic              float64 = 0.53
	volSteps           int     = 20
	stepSize           float64 = 0.1
	zoom               float64 = 0.800
	tile               float64 = 0.850
	speed              float64 = 0.010
	brightness         float64 = 0.0015
	darkMatter         float64 = 0.300
	distFading         float64 = 0.730
	saturation         float64 = 0.850
	frequencyVariation float64 = 1.3
	sparsity           float64 = 0.5
)

var cos = math.Cos
var sin = math.Sin

// https://www.shadertoy.com/view/XlfGRj
func Cosmic1(img *image.RGBA) {
	size := img.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			// coords & direction
			uv := m.Vec2{float64(x)/float64(size.X) - 0.5, float64(y)/float64(size.Y) - 0.5}
			uv = m.Vec2{uv.X(), uv.Y() * float64(size.Y) / float64(size.X)}
			dir := m.Vec3{uv[0] * zoom, uv[1] * zoom, 1.0}
			from := m.Vec3{1, 0.5, 0.5}

			// volumetric rendering
			s := 0.1
			fade := 1.0
			v := m.Vec3{}
			for r := 0; r < volSteps; r++ {
				var p m.Vec3
				p = from.Add(m.Vec3{s, s, s})
				p = m.Vec3{p.X() * dir.X(), p.Y() * dir.Y(), p.Z() * dir.Z()}.Mul(.5)
				p = m.Vec3{tile, 0, 0}.Sub(m.Vec3{math.Mod(p.X(), tile*2), p.Y(), p.Z()})
				p = m.Vec3{math.Abs(p.X()), math.Abs(p.Y()), math.Abs(p.Z())}
				var a, pa float64
				for i := 0; i < iterations; i++ {
					mv := p.Dot(p) - magic
					p = m.Vec3{p.X() / mv, p.Y() / mv, p.Z() / mv}
					a += math.Abs(p.Len() - pa)
					pa = p.Len()
				}
				dm := math.Max(0, darkMatter-a*a*.001)
				a *= a * a
				if r > 6 {
					fade *= 1 - dm
				}
				v = m.Vec3{v.X() + fade, v.Y() + fade, v.Z() + fade}
				abf := a * brightness * fade
				v = m.Vec3{v.X() + s*abf, v.Y() + s*s*abf, v.Z() + s*s*s*s*abf}
				fade *= distFading
				s += stepSize
			}

			//color adjust
			z := v.Len()
			v = v.Add(m.Vec3{
				lerp(z, v.X(), saturation),
				lerp(z, v.Y(), saturation),
				lerp(z, v.Z(), saturation)})
			c := m.Vec4{v.X() * .01, v.Y() * .01, v.Z() * 01, 1}
			color := color.RGBA{toUint8(c.X()), toUint8(c.Y()), toUint8(c.Z()), toUint8(c.W())}

			img.SetRGBA(x, y, color)
		}
	}
}

func toUint8(in float64) uint8 {
	v := int(in * 256)
	if v >= 256 {
		return 255
	} else if v <= 0 {
		return 0
	}
	return uint8(v)
}

func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

//  https://casual-effects.blogspot.com/2013/08/starfield-shader.html
//  problem: output image is always the same mono color
func Cosmic2(img *image.RGBA) {
	size := img.Bounds().Size()
	var resolution = m.Vec2{float64(size.X), float64(size.Y)}
	var invResolution = m.Vec2{resolution.Y(), resolution.X()}

	rrt := float64(mr.Intn(size.X))
	rt := .5 + (rrt/3)/float64(size.X)*2
	var rotate = m.Mat2{math.Cos(rt), math.Sin(rt), -(math.Sin(rt)), math.Cos(rt)}

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			uv := m.Vec2{float64(size.X)*resolution.X() - 0.5, float64(size.Y)*resolution.Y() - 0.5*(resolution.Y()*invResolution.X())}
			dir := m.Vec3{(uv.X() * zoom) * rotate.At(0, 0), uv.Y() * zoom, 1.0 * rotate.At(1, 0)}
			s := 0.1
			fade := 0.01
			var c m.Vec3
			var o m.Vec3
			for r := 0; r < volSteps; r++ {
				pp := m.Vec3{float64(x) + dir.X()*(s*0.5), float64(y) + dir.Y()*(s*0.5), dir.Z() * (s * 0.5)}
				fv := func(in float64) float64 {
					return math.Abs(frequencyVariation - math.Mod(in, (frequencyVariation*2)))
				}
				p := m.Vec3{fv(pp.X()), fv(pp.Y()), fv(pp.Z())}
				prevLen := 0.0
				a := 0.0
				for i := 0; i < iterations; i++ {
					p = m.Vec3{math.Abs(p.X()), math.Abs(p.Y()), math.Abs(p.Z())}
					p = p.Mul(1.0/p.Dot(p) + (-sparsity))
					l := p.Len()
					a += math.Abs(l - prevLen)
					prevLen = l
				}
				a *= a * a
				tw := m.Vec3{s, s * s, s * s * s}
				tw.Mul((a*brightness + 1.0))
				tw.Mul(fade)
				c = m.Vec3{c.X() + tw.X(), c.Y() + tw.Y(), c.Z() + tw.Z()}
				fade *= distFading
				s += stepSize
			}
			c = m.Vec3{math.Min(c.X(), 1.2), c.Y(), c.Z()}
			intensity := math.Min((c.X() + c.Y() + c.Z()), 0.7)
			sgn := m.Vec2{float64(x) + 1.0, float64(y) + 1.0}
			sgn.Mul(2.0)
			sgn.Sub(m.Vec2{1, 1})
			gradient := m.Vec2{intensity * sgn.X(), intensity * sgn.Y()}
			cutoff := math.Max(math.Max(gradient.X(), gradient.Y())-0.1, 0.0)
			c.Mul(math.Max(1.0-cutoff*6.0, 0.3))
			c = m.Vec3{lerp(c.X(), o.X(), intensity) - .004, lerp(c.Y(), o.Y(), intensity), lerp(c.Z(), o.Z(), intensity)}
			o = c

			clr := color.RGBA{toUint8(c.X()), toUint8(c.Y()), toUint8(c.Z()), toUint8(1)}
			img.SetRGBA(x, y, clr)
		}
	}
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	Cosmic1(m)
	f, err := os.Create("main.png")
	check(err)
	check(png.Encode(f, m))
	check(f.Close())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
