package main

import "math"

type RGBA uint32
type RGB struct{ R, G, B uint8 }
type HSL struct{ H, S, L float64 }

func (rgba RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(rgba >> 24 & 0xFF)
	g = uint32(rgba >> 16 & 0xFF)
	b = uint32(rgba >> 8 & 0xFF)
	a = uint32(rgba >> 0 & 0xFF)
	r |= r << 8
	g |= g << 8
	b |= b << 8
	a |= a << 8
	return
}

func (rgb RGB) RGBA() (r, g, b, a uint32) {
	r, g, b = uint32(rgb.R), uint32(rgb.G), uint32(rgb.B)
	r |= r << 8
	g |= g << 8
	b |= b << 8
	a = 0xFFFF
	return
}

func (hsl HSL) RGBA() (r, g, b, a uint32) {
	r1, g1, b1, _ := hsla(hsl.H, hsl.S, hsl.L, 1)
	return sat16(r1), sat16(g1), sat16(b1), 0xFFFF
}

func hue(v1, v2, h float64) float64 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	if 6*h < 1 {
		return v1 + (v2-v1)*6*h
	} else if 2*h < 1 {
		return v2
	} else if 3*h < 2 {
		return v1 + (v2-v1)*(2.0/3.0-h)*6
	}

	return v1
}

func hsla(h, s, l, a float64) (r, g, b, ra float64) {
	if s == 0 {
		return l, l, l, a
	}

	h = math.Mod(h, TAU) / TAU

	var v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - s*l
	}

	v1 := 2*l - v2
	r = hue(v1, v2, h+1.0/3.0)
	g = hue(v1, v2, h)
	b = hue(v1, v2, h-1.0/3.0)
	ra = a

	return
}

// sat16 converts 0..1 float to 0..0xFFFF uint32
func sat16(v float64) uint32 {
	v = v * 0xFFFF
	if v >= 0xFFFF {
		return 0xFFFF
	} else if v <= 0 {
		return 0
	}
	return uint32(v) & 0xFFFF
}
