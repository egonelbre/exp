package mldsa

import (
	"math/rand/v2"
	"testing"
)

func TestNTTMulUnroll(t *testing.T) {
	for range 10 {
		var a, b NTTElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
			b[i] = FieldElement(rand.Uint32())
		}

		exp := NTTMul(a, b)
		got := NTTMulUnroll(a, b)

		if exp != got {
			t.Fatal("wrong result")
		}
	}
}

func TestInverseNTTUnroll(t *testing.T) {
	for range 10 {
		var a NTTElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := InverseNTT(a)
		got := InverseNTTUnroll(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}

func TestNTTUnroll(t *testing.T) {
	for range 10 {
		var a RingElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := NTT(a)
		got := NTTUnroll(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}
