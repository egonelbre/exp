package main

import "image/color"

type HL struct{ H, L float32 }

type Image struct {
	Width  int
	Height int
	Pix    []HL
}

func NewImage() *Image {
	return &Image{}
}

func (m *Image) SetSize(w, h int) bool {
	if m.Width != w || m.Height != h {
		m.Width = w
		m.Height = h
		m.Pix = make([]HL, w*h)
		return true
	}
	return false
}

func (m *Image) Clear() {
	for i := range m.Pix {
		m.Pix[i] = HL{}
	}
}

func (m *Image) Put(v V, b HL) {
	x, y := int(v.X), int(v.Y)
	if x < 0 || m.Width <= x {
		return
	}
	if y < 0 || m.Height <= y {
		return
	}
	i := y*m.Width + x
	m.Pix[i] = blendhl(m.Pix[i], b)
}

func blendhl(a, b HL) HL {
	return HL{
		H: (a.H*a.L + b.H*b.L) / (a.L + b.L),
		L: a.L + b.L,
	}
}

func (hl HL) RGBA(scale float32) color.RGBA {
	r, g, b, _ := hsla(hl.H, 0.7, Clamp1(hl.L*scale), 1)
	return color.RGBA{sat8(r), sat8(g), sat8(b), 0xFF}
}
