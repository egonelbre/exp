package vector

import (
	"runtime"
	"testing"
)

type Vec3 [3]float32

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}
func (a Vec3) Mul(v float32) Vec3 {
	return Vec3{a[0] * v, a[1] * v, a[2] * v}
}

func BenchmarkAddMul(b *testing.B) {
	x, y := Vec3{1, 2, 3}, Vec3{5, 6, 7}
	for i := 0; i < b.N; i++ {
		x = x.Add(y.Mul(0.15))
	}
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

func BenchmarkAddMulFused(b *testing.B) {
	x, y := Vec3{1, 2, 3}, Vec3{5, 6, 7}
	for i := 0; i < b.N; i++ {
		x[0] += y[0] * 0.15
		x[1] += y[1] * 0.15
		x[2] += y[2] * 0.15
	}
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

type Vec3s struct{ X, Y, Z float32 }

func (a Vec3s) Add(b Vec3s) Vec3s {
	return Vec3s{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}
func (a Vec3s) Mul(v float32) Vec3s {
	return Vec3s{a.X * v, a.Y * v, a.Z * v}
}

func BenchmarkStructAddMul(b *testing.B) {
	x, y := Vec3s{1, 2, 3}, Vec3s{5, 6, 7}
	for i := 0; i < b.N; i++ {
		x = x.Add(y.Mul(0.15))
	}
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

func BenchmarkStructAddMulFused(b *testing.B) {
	x, y := Vec3s{1, 2, 3}, Vec3s{5, 6, 7}
	for i := 0; i < b.N; i++ {
		x.X += y.X * 0.15
		x.Y += y.Y * 0.15
		x.Z += y.Z * 0.15
	}
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}
