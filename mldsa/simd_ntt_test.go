package mldsa

import (
	"simd"
	"testing"
)

func TestVecMontgomeryMul(t *testing.T) {
	var aArr, bArr [8]uint32
	for i := 0; i < 8; i++ {
		aArr[i] = uint32(i + 100)
		bArr[i] = uint32(i + 200)
	}

	aVec := simd.LoadUint32x8(&aArr)
	bVec := simd.LoadUint32x8(&bArr)

	resVec := vecMontgomeryMul(aVec, bVec)

	var resArr [8]uint32
	resVec.Store(&resArr)

	for i := 0; i < 8; i++ {
		expected := FieldMontgomeryMul(FieldElement(aArr[i]), FieldElement(bArr[i]))
		if uint32(expected) != resArr[i] {
			t.Errorf("idx %d: expected %d, got %d", i, expected, resArr[i])
		}
	}
}

func TestNTTMul_AVX2(t *testing.T) {
	var a, b NTTElement
	for i := 0; i < 256; i++ {
		a[i] = FieldElement(i)
		b[i] = FieldElement(i + 1)
	}

	expected := NTTMul(a, b)
	got := NTTMul_AVX2(&a, &b)

	for i := 0; i < 256; i++ {
		if expected[i] != got[i] {
			t.Errorf("idx %d: expected %d, got %d", i, expected[i], got[i])
		}
	}
}

func TestNTT_AVX2(t *testing.T) {
	var f NTTElement
	for i := 0; i < 256; i++ {
		f[i] = FieldElement(i)
	}

	expected := NTT(RingElement(f))

	got := f
	NTT_AVX2(&got)

	for i := 0; i < 256; i++ {
		if expected[i] != got[i] {
			t.Errorf("idx %d: expected %d, got %d", i, expected[i], got[i])
		}
	}
}

func TestInverseNTT_AVX2(t *testing.T) {
	var f NTTElement
	for i := 0; i < 256; i++ {
		f[i] = FieldElement(i)
	}

	// Create a valid NTT domain input
	nttF := NTT(RingElement(f))

	expected := InverseNTT(nttF)

	got := nttF
	InverseNTT_AVX2(&got)

	for i := 0; i < 256; i++ {
		if expected[i] != got[i] {
			t.Errorf("idx %d: expected %d, got %d", i, expected[i], got[i])
		}
	}
}
