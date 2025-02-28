package main

import "testing"

func mula(u, v, x, y, z, w  uint64) uint64 {
    return u * v * x * y * z * w
}

func mulb(u, v, x, y, z, w uint64) uint64 {
	a := u * v
    b := x * y
    c := z * w
    return a * b * c
}

var u, v, x, y, z, w, r uint64

func BenchmarkA(b *testing.B) {
	for b.Loop() {
		r += mula(u, v, x, y, z, w)
	}
}

func BenchmarkB(b *testing.B) {
	for b.Loop() {
		r += mulb(u, v, x, y, z, w)
	}
}
func BenchmarkA1(b *testing.B) {
	for b.Loop() {
		r += mula(u, v, x, y, z, w)
	}
}

func BenchmarkB1(b *testing.B) {
	for b.Loop() {
		r += mulb(u, v, x, y, z, w)
	}
}
func BenchmarkA2(b *testing.B) {
	for b.Loop() {
		r += mula(u, v, x, y, z, w)
	}
}

func BenchmarkB2(b *testing.B) {
	for b.Loop() {
		r += mulb(u, v, x, y, z, w)
	}
}