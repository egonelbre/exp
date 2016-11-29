package ui

type Rect struct {
	Left, Top     int
	Width, Height int
}

func (r Rect) Bounds() Rect { return r }
