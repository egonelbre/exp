package main

import (
	"bytes"
	"fmt"
)

type Unit struct {
	meter    int8
	kilogram int8
	second   int8
	ampere   int8
	kelvin   int8
	mole     int8
	candela  int8
}

var (
	Scalar = Unit{0, 0, 0, 0, 0, 0, 0}
	M      = Unit{1, 0, 0, 0, 0, 0, 0}
	Kg     = Unit{0, 1, 0, 0, 0, 0, 0}
	Sec    = Unit{0, 0, 1, 0, 0, 0, 0}

	MetersPerSec = Unit{1, 0, -1, 0, 0, 0, 0}
)

func (a *Unit) Mul(b Unit) {
	a.meter += b.meter
	a.kilogram += b.kilogram
	a.second += b.second
	a.ampere += b.ampere
	a.kelvin += b.kelvin
	a.mole += b.mole
	a.candela += b.candela
}

func (a *Unit) Div(b Unit) {
	a.meter -= b.meter
	a.kilogram -= b.kilogram
	a.second -= b.second
	a.ampere -= b.ampere
	a.kelvin -= b.kelvin
	a.mole -= b.mole
	a.candela -= b.candela
}

func (u *Unit) String() string {
	var b bytes.Buffer
	a := func(prefix string, v int8) {
		if v != 0 {
			b.WriteString(fmt.Sprintf("%s^%d", prefix, v))
		}
	}
	a("m", u.meter)
	a("kg", u.kilogram)
	a("s", u.second)

	a("A", u.ampere)
	a("K", u.kelvin)
	a("mol", u.mole)
	a("cd", u.candela)
	return b.String()
}

type V struct {
	Value float64
	Unit  Unit
}

func (a V) Add(b V) V {
	if a.Unit != b.Unit {
		panic("Incompatible units")
	}
	r := a
	r.Value += b.Value
	return r
}

func (a V) Sub(b V) V {
	if a.Unit != b.Unit {
		panic("Incompatible units")
	}
	r := a
	r.Value -= b.Value
	return r
}

func (a V) Mul(b V) V {
	r := a
	r.Unit.Mul(b.Unit)
	r.Value *= b.Value
	return r
}

func (a V) Div(b V) V {
	r := a
	r.Unit.Div(b.Unit)
	r.Value /= b.Value
	return r
}

func (a V) String() string {
	return fmt.Sprintf("%f[%s]", a.Value, a.Unit.String())
}

func main() {
	a := V{3, Kg}
	a.Unit.kilogram = 1

	b := V{2, Sec}

	r := a.Div(b)
	fmt.Printf("%v\n", r)
	r = r.Div(b)
	fmt.Printf("%v\n", r)

	v := r.Add(b)
	fmt.Printf("%v\n", v)
}
