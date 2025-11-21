package mldsa

import (
	"math/rand/v2"
	"testing"
)

func TestInverseNTTUnroll2(t *testing.T) {
	for range 10 {
		var a NTTElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := InverseNTT(a)
		got := InverseNTTUnroll2(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}

func TestInverseNTTUnroll3(t *testing.T) {
	for range 10 {
		var a NTTElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := InverseNTT(a)
		got := InverseNTTUnroll3(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}

func TestNTTUnroll2(t *testing.T) {
	for range 10 {
		var a RingElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := NTT(a)
		got := NTTUnroll2(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}
