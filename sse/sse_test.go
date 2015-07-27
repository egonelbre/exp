package sse

import (
	"testing"
	"testing/quick"
)

func same(a, b []uint32) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type op func(a, b []uint32)

func testOp(t *testing.T, fast, slow op) {
	f := func(dst1, src []uint32) bool {
		dst2 := make([]uint32, len(dst1))
		copy(dst2, dst1)

		slow(dst2, src)
		fast(dst1, src)

		return same(dst1, dst2)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMinU32Len(t *testing.T) {
	var d, s [16]uint32
	if v := MinU32Len(d[:7], s[:3]); v != 3 {
		t.Errorf("1. got %v expected %v", v, 3)
	}

	if v := MinU32Len(d[:3], s[:7]); v != 3 {
		t.Errorf("2. got %v expected %v", v, 3)
	}

	if v := MinU32Len(d[:7], s[:7]); v != 7 {
		t.Errorf("3. got %v expected %v", v, 7)
	}

	if v := MinU32Len(d[:0], s[:0]); v != 0 {
		t.Errorf("4. got %v expected %v", v, 0)
	}
}

func TestAddU32(t *testing.T) {
	dst := []uint32{0, 1, 2, 3, 4}
	src := []uint32{4, 4, 2, 1}
	exp := []uint32{4, 5, 4, 4, 4}
	AddU32(dst, src)
	if !same(dst, exp) {
		t.Errorf("got %v exp %v", dst, exp)
	}
}

func TestAddU32_Quick(t *testing.T) { testOp(t, AddU32, addU32) }

func TestSubU32(t *testing.T) {
	dst := []uint32{7, 9, 2, 3, 4}
	src := []uint32{4, 4, 6, 9}
	exp := []uint32{3, 5, ^uint32(4) + 1, ^uint32(6) + 1, 4}
	SubU32(dst, src)
	if !same(dst, exp) {
		t.Errorf("got %v exp %v", dst, exp)
	}
}
func TestSubU32_Quick(t *testing.T) { testOp(t, SubU32, subU32) }
func TestMulU32_Quick(t *testing.T) { testOp(t, MulU32, mulU32) }
