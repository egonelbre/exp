package g

type Vector struct{ X, Y float32 }

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}
