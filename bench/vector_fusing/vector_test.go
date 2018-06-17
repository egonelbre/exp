package vector

import "testing"

type Vec3 [3]float32

func Add(a, b Vec3) Vec3 {
	return Vec3{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}
func Mul(a Vec3, v float32) Vec3 {
	return Vec3{a[0] * v, a[1] * v, a[2] * v}
}

func BenchmarkAddMul(b *testing.B) {
	x, y := Vec3{1, 2, 3}, Vec3{5, 6, 7}
	for i := 0; i < b.N; i++ {
		x = Add(x, Mul(y, 0.15))
	}
}

func BenchmarkAddMulFused(b *testing.B) {
	x, y := Vec3{1, 2, 3}, Vec3{5, 6, 7}
	for i := 0; i < b.N; i++ {
		x[0] += y[0] * 0.15
		x[1] += y[1] * 0.15
		x[2] += y[2] * 0.15
	}
}
