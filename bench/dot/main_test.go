package main

import (
	"testing"
	"unsafe"
)

func Dot1(a, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i += 4 {
		s0 := a[i] * b[i]
		s1 := a[i+1] * b[i+1]
		s2 := a[i+2] * b[i+2]
		s3 := a[i+3] * b[i+3]
		sum += s0 + s1 + s2 + s3
	}
	return sum
}

func Dot2(a, b []float64) float64 {
	type F64x4 struct{ X, Y, Z, W float64 }

	xs := unsafe.Slice((*F64x4)(unsafe.Pointer(unsafe.SliceData(a))), len(a)/4)
	ys := unsafe.Slice((*F64x4)(unsafe.Pointer(unsafe.SliceData(b))), len(b)/4)

	if len(xs) != len(ys) {
		return 0
	}

	var x, y, z, w float64

	for i := range xs {
		x += xs[i].X * ys[i].X
		y += xs[i].Y * ys[i].Y
		z += xs[i].Z * ys[i].Z
		w += xs[i].W * ys[i].W
	}

	// TODO: handle leftover stuff

	return x + y + z + w
}

func Dot3(a, b []float64) float64 {
	var x, y, z, w float64

	for i := 0; i < len(a); i += 4 {
		x += a[i] * b[i]
		y += a[i+1] * b[i+1]
		z += a[i+2] * b[i+2]
		w += a[i+3] * b[i+3]
	}

	return x + y + z + w
}

func Dot4(a, b []float64) float64 {
	var sum float64

	for len(a) >= 4 && len(b) >= 4 {
		s0 := a[0] * b[0]
		s1 := a[1] * b[1]
		s2 := a[2] * b[2]
		s3 := a[3] * b[3]
		sum += s0 + s1 + s2 + s3
		a = a[4:]
		b = b[4:]
	}
	// deal with remainder of a and b
	return sum
}

func Dot5(a, b []float64) float64 {
	var x, y, z, w float64

	for len(a) >= 4 && len(b) >= 4 {
		x += a[0] * b[0]
		y += a[1] * b[1]
		z += a[2] * b[2]
		w += a[3] * b[3]
		a, b = a[4:], b[4:]
	}

	return x + y + z + w
}

var (
	xs = make([]float64, 10*1024)
	ys = make([]float64, 10*1024)

	sink float64
)

func BenchmarkDot1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot1(xs, ys)
	}
}

func BenchmarkDot2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot2(xs, ys)
	}
}

func BenchmarkDot3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot3(xs, ys)
	}
}

func BenchmarkDot4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot4(xs, ys)
	}
}

func BenchmarkDot5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot5(xs, ys)
	}
}
