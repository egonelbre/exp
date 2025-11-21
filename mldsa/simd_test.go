package mldsa

import (
	"math/rand/v2"
	simd "simd/archsimd"
	"testing"
	"unsafe"
)

func FuzzFieldReduceOnce8(f *testing.F) {
	var z uint32
	f.Add(z, z, z, z, z, z, z, z)
	f.Fuzz(func(t *testing.T, a0, a1, a2, a3, a4, a5, a6, a7 uint32) {
		in := [8]uint32{a0, a1, a2, a3, a4, a5, a6, a7}
		exp := [8]FieldElement{}

		for i := range in {
			exp[i] = FieldReduceOnce(in[i])
		}

		in8 := Uint32x8FromArray(in)
		got8 := FieldReduceOnce8(in8)
		got := FieldElement8AsArray(got8)

		if exp != got {
			t.Fatalf("wrong result in:%v exp:%v got:%v", in, exp, got)
		}
	})
}

func FuzzFieldReduceOnceRef(f *testing.F) {
	var z uint32
	f.Add(z)
	f.Fuzz(func(t *testing.T, a0 uint32) {
		exp := FieldReduceOnce(a0)
		got := FieldReduceOnceRef(a0)
		if exp != got {
			t.Fatalf("wrong result a:%v exp:%v got:%v", a0, exp, got)
		}
	})
}

func FuzzFieldAdd8(f *testing.F) {
	var z uint32
	f.Add(
		z, z, z, z, z, z, z, z,
		z, z, z, z, z, z, z, z,
	)
	f.Fuzz(func(t *testing.T,
		a0, a1, a2, a3, a4, a5, a6, a7 uint32,
		b0, b1, b2, b3, b4, b5, b6, b7 uint32,
	) {
		inA := [8]uint32{a0, a1, a2, a3, a4, a5, a6, a7}
		inB := [8]uint32{b0, b1, b2, b3, b4, b5, b6, b7}
		exp := [8]FieldElement{}

		for i := range inA {
			exp[i] = FieldAdd(FieldElement(inA[i]), FieldElement(inB[i]))
		}

		inA8 := Uint32x8FromArray(inA)
		inB8 := Uint32x8FromArray(inB)
		got8 := FieldAdd8(inA8, inB8)
		got := FieldElement8AsArray(got8)

		if exp != got {
			t.Fatalf("wrong result a:%v b:%v exp:%v got:%v", inA, inB, exp, got)
		}
	})
}

func FuzzFieldSub8(f *testing.F) {
	var z uint32
	f.Add(
		z, z, z, z, z, z, z, z,
		z, z, z, z, z, z, z, z,
	)
	f.Fuzz(func(t *testing.T,
		a0, a1, a2, a3, a4, a5, a6, a7 uint32,
		b0, b1, b2, b3, b4, b5, b6, b7 uint32,
	) {
		inA := [8]uint32{a0, a1, a2, a3, a4, a5, a6, a7}
		inB := [8]uint32{b0, b1, b2, b3, b4, b5, b6, b7}
		exp := [8]FieldElement{}

		for i := range inA {
			exp[i] = FieldSub(FieldElement(inA[i]), FieldElement(inB[i]))
		}

		inA8 := Uint32x8FromArray(inA)
		inB8 := Uint32x8FromArray(inB)
		got8 := FieldSub8(inA8, inB8)
		got := FieldElement8AsArray(got8)

		if exp != got {
			t.Fatalf("wrong result a:%v b:%v exp:%v got:%v", inA, inB, exp, got)
		}
	})
}

func FuzzFieldMontgomeryMul8(f *testing.F) {
	var z uint32
	f.Add(
		z, z, z, z, z, z, z, z,
		z, z, z, z, z, z, z, z,
	)
	f.Fuzz(func(t *testing.T,
		a0, a1, a2, a3, a4, a5, a6, a7 uint32,
		b0, b1, b2, b3, b4, b5, b6, b7 uint32,
	) {
		inA := [8]uint32{a0, a1, a2, a3, a4, a5, a6, a7}
		inB := [8]uint32{b0, b1, b2, b3, b4, b5, b6, b7}
		exp := [8]FieldElement{}

		for i := range inA {
			exp[i] = FieldMontgomeryMul(FieldElement(inA[i]), FieldElement(inB[i]))
		}

		inA8 := Uint32x8FromArray(inA)
		inB8 := Uint32x8FromArray(inB)
		got8 := FieldMontgomeryMul8(inA8, inB8)
		got := FieldElement8AsArray(got8)

		if exp != got {
			t.Fatalf("wrong result a:%v b:%v exp:%v got:%v", inA, inB, exp, got)
		}
	})
}

func TestNTTMul8(t *testing.T) {
	for range 10 {
		var a, b NTTElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
			b[i] = FieldElement(rand.Uint32())
		}

		exp := NTTMul(a, b)
		got := NTTMul8(a, b)

		if exp != got {
			t.Fatal("wrong result")
		}
	}
}

func TestInverseNTT8(t *testing.T) {
	for range 10 {
		var a NTTElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := InverseNTT(a)
		got := InverseNTT8(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}

func TestNTT8(t *testing.T) {
	for range 10 {
		var a RingElement
		for i := range a {
			a[i] = FieldElement(rand.Uint32())
		}

		exp := NTT(a)
		got := NTT8(a)

		if exp != got {
			t.Fatalf("wrong result\nin:%v\nexp:%v\ngot:%v", a, exp, got)
		}
	}
}

func FieldElement8FromArray(x [8]FieldElement) FieldElement8 {
	ptr := (*[8]uint32)(unsafe.Pointer(&x))
	return simd.LoadUint32x8(ptr)
}

func Uint32x8FromArray(x [8]uint32) simd.Uint32x8 {
	ptr := (*[8]uint32)(unsafe.Pointer(&x[0]))
	return simd.LoadUint32x8(ptr)
}

func FieldElement8AsArray(x FieldElement8) [8]FieldElement {
	var res [8]uint32
	x.StoreSlice(res[:])
	return *(*[8]FieldElement)(unsafe.Pointer(&res[0]))
}
