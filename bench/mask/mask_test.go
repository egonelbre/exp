package main

import "testing"

func MaskAndConst()

func BenchmarkMaskAndConst(b *testing.B) {
	for b.Loop() {
		MaskAndConst()
	}
}

func MaskShlShr()
func BenchmarkMaskShlShr(b *testing.B) {
	for b.Loop() {
		MaskShlShr()
	}
}

func MaskAndAddr()
func BenchmarkMaskAndAddr(b *testing.B) {
	for b.Loop() {
		MaskAndAddr()
	}
}
