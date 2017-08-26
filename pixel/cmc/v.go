package main

type V struct{ X, Y, Z float32 }

func S(v V, s float32) V { return V{v.X * s, v.Y * s, v.Z * s} }
func L(v V) float32      { return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z) }
func N(v V, s float32) V { return S(v, s/L(v)) }

func Rx(v V, s float32) V {
	sn, cs := Sincos(s)
	return V{v.X, v.Y*cs - v.Z*sn, v.Y*sn + v.Z*cs}
}

func Ry(v V, s float32) V {
	sn, cs := Sincos(s)
	return V{v.Z*cs - v.X*sn, v.Y, v.Z*sn + v.X*cs}
}

func P(v, C V) (r V) {
	r.Z = (v.Z*0.5 + 30) / 30
	r.X = C.X - v.X*r.Z*10
	r.Y = C.Y - v.Y*r.Z*10
	r.Z = C.Z + r.Z*3
	return r
}
