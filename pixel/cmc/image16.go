package main

type Image16 struct {
	Width  int
	Height int
	Pix    []uint16
}

func NewImage16() *Image16 {
	return &Image16{}
}

func (m *Image16) SetSize(w, h int) bool {
	if m.Width != w || m.Height != h {
		m.Width = w
		m.Height = h
		m.Pix = make([]uint16, w*h)
		return true
	}
	return false
}

func (m *Image16) Clear() {
	for i := range m.Pix {
		m.Pix[i] = 0
	}
}

func (m *Image16) Put(v V) {
	x, y := int(v.X), int(v.Y)
	if x < 0 || m.Width <= x {
		return
	}
	if y < 0 || m.Height <= y {
		return
	}

	off := y*m.Width + x
	if m.Pix[off] < 0xFFFF {
		m.Pix[off]++
	}
}
