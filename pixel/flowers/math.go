package main

import (
	"math"

	"github.com/faiface/pixel"
)

const TAU = 2 * math.Pi

func Bounce(t, start, end, duration float64) float64 {
	t = t / duration
	ts := t * t
	tc := ts * t
	return start + end*(33*tc*ts+-106*ts*ts+126*tc+-67*ts+15*t)
}

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
