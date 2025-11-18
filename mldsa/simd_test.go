package mldsa

import (
	"math/rand/v2"
	"simd"
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

func FuzzFieldMontgomeryReduce4(f *testing.F) {
	var z uint64
	f.Add(z, z, z, z)
	f.Fuzz(func(t *testing.T, a0, a1, a2, a3 uint64) {
		in := [4]uint64{a0, a1, a2, a3}
		exp := [4]FieldElement{}

		for i := range in {
			exp[i] = FieldMontgomeryReduce(in[i])
		}

		in8 := Uint64x4FromArray(in)
		got8 := FieldMontgomeryReduce4(in8)
		got := FieldElement4xAsArray(got8)

		if exp != got {
			t.Fatalf("wrong result in:%v exp:%v got:%v", in, exp, got)
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

func TestInverseNTT8(t *testing.T) {
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

func TestNTT8(t *testing.T) {
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

func TestUint64x4_ConvertToUint32(t *testing.T) {
	a := [4]uint64{1, 2, 3, 0xFFFF_FFFF_FFFF_FFFF}
	reg64 := simd.LoadUint64x4(&a)
	reg32 := Uint64x4_ConvertToUint32(reg64)
	var b [4]uint32
	reg32.Store(&b)
	if b != [4]uint32{1, 2, 3, 0xFFFF_FFFF} {
		t.Fatalf("got %v", b)
	}
}

func TestUint32x4_Spread(t *testing.T) {
	a := [4]uint32{1, 2, 3, 4}
	reg := simd.LoadUint32x4(&a)
	spread := Uint32x4_Spread(reg)
	var b [8]uint32
	spread.Store(&b)
	exp := [8]uint32{0, 1, 0, 2, 0, 3, 0, 4}
	if b != exp {
		t.Fatalf("got %v, want %v", b, exp)
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

func Uint64x4FromArray(x [4]uint64) simd.Uint64x4 {
	ptr := (*[4]uint64)(unsafe.Pointer(&x[0]))
	return simd.LoadUint64x4(ptr)
}

func FieldElement8AsArray(x FieldElement8) [8]FieldElement {
	var res [8]uint32
	x.StoreSlice(res[:])
	return *(*[8]FieldElement)(unsafe.Pointer(&res[0]))
}

func FieldElement4xAsArray(x FieldElement4) [4]FieldElement {
	var res [4]uint32
	x.StoreSlice(res[:])
	return *(*[4]FieldElement)(unsafe.Pointer(&res[0]))
}
