package g

type Vec2 struct { X, Y float64 }

// Add adds two vectors component-wise.
//
//go:fix inline
func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{X: v.X + u.X, Y: v.Y + u.Y}
}

// Scale multiples each component by u.
//
//go:fix inline
func (v Vec2) Scale(u float64) Vec2 {
	return Vec2{X: v.X * u, Y: v.Y * u}
}
