package main

import "testing"

type Large [4 * 1024]byte

const N = 1024

func IndexA(s []Large, v Large) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func IndexB(s []Large, v Large) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func BenchmarkA_Large(b *testing.B) {
	elements := make([]Large, N)
	for k := 0; k < b.N; k++ {
		IndexA(elements, Large{1})
	}
}

func BenchmarkB_Large(b *testing.B) {
	elements := make([]Large, N)
	for k := 0; k < b.N; k++ {
		IndexB(elements, Large{1})
	}
}