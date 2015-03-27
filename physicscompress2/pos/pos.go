package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func sin(v float32) float32  { return float32(math.Sin(float64(v))) }
func cos(v float32) float32  { return float32(math.Cos(float64(v))) }
func sqrt(v float32) float32 { return float32(math.Sqrt(float64(v))) }

func atan2(x, y float32) float32 { return float32(math.Atan2(float64(x), float64(y))) }
func acos(x float32) float32     { return float32(math.Acos(float64(x))) }

type XYZ struct{ X, Y, Z float32 }

func (p XYZ) Polar() Polar {
	r := p.Len()
	return Polar{
		R: r,
		O: atan2(p.Y, p.X),
		F: acos(p.Z / r),
	}
}

func (a XYZ) Sub(b XYZ) XYZ {
	return XYZ{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (p XYZ) Len() float32 {
	return sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p XYZ) Norm() XYZ         { return p.Div(p.Len()) }
func (p XYZ) Div(v float32) XYZ { return XYZ{X: p.X / v, Y: p.Y / v, Z: p.Z / v} }
func (p XYZ) Mul(v float32) XYZ { return XYZ{X: p.X * v, Y: p.Y * v, Z: p.Z * v} }

func (p XYZ) Round(t float32) XYZ {
	return XYZ{
		X: float32(int(p.X*t)) / t,
		Y: float32(int(p.Y*t)) / t,
		Z: float32(int(p.Z*t)) / t,
	}
}

type Polar struct{ R, F, O float32 }

func (p Polar) XYZ() XYZ {
	return XYZ{
		X: p.R * cos(p.O) * sin(p.F),
		Y: p.R * sin(p.O) * sin(p.F),
		Z: p.R * cos(p.F),
	}
}

func (p Polar) Round(t float32) Polar {
	return Polar{
		R: float32(int(p.R*t)) / t,
		O: float32(int(p.O*t)) / t,
		F: float32(int(p.F*t)) / t,
	}
}

func Rand(s float32) XYZ {
	return XYZ{
		X: rand.Float32()*s - s/2,
		Y: rand.Float32()*s - s/2,
		Z: rand.Float32()*s - s/2,
	}
}

func RandLen(x float32) XYZ {
	return XYZ{
		X: rand.Float32() - 1/2,
		Y: rand.Float32() - 1/2,
		Z: rand.Float32() - 1/2,
	}.Norm().Mul(x)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := RandLen(40.0)
	for i := 0; i < 8; i += 1 {
		t := a.Polar().Round(float32(uint(1) << uint(i)))
		b := t.XYZ()
		d := a.Sub(b)
		l := d.Len()

		fmt.Println(i)
		fmt.Println("  L:", l)
		fmt.Println("  S:", a)
		fmt.Println(" S2:", a.Round(float32(uint(1)<<uint(i))))
		fmt.Println("  E:", b)
		fmt.Println("  D:", d)
	}

}
