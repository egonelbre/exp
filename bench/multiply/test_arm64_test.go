package main

import "testing"

func Mul19()
func Mul19shift()

func BenchmarkMul19(b *testing.B) {
	for b.Loop() {
		Mul19()
	}
}

func BenchmarkMul19shift(b *testing.B) {
	for b.Loop() {
		Mul19shift()
	}
}

func ManyMul19()
func ManyMul19shift()

func BenchmarkManyMul19(b *testing.B) {
	for b.Loop() {
		ManyMul19()
	}
}

func BenchmarkManyMul19shift(b *testing.B) {
	for b.Loop() {
		ManyMul19shift()
	}
}
