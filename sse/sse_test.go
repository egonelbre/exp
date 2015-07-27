package sse

import (
	"math/rand"
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

func TestAddU32_ASM(t *testing.T) {
	dst := []uint32{0, 1, 2, 3, 4}
	src := []uint32{4, 4, 2, 1}
	exp := []uint32{4, 5, 4, 4, 4}
	AddU32_ASM(dst, src)
	if !same(dst, exp) {
		t.Errorf("got %v exp %v", dst, exp)
	}
}

func TestSubU32_ASM(t *testing.T) {
	dst := []uint32{7, 9, 2, 3, 4}
	src := []uint32{4, 4, 6, 9}
	exp := []uint32{3, 5, ^uint32(4) + 1, ^uint32(6) + 1, 4}
	SubU32_ASM(dst, src)
	if !same(dst, exp) {
		t.Errorf("got %v exp %v", dst, exp)
	}
}

func TestAddU32_ASM_Quick(t *testing.T) { testOp(t, AddU32_ASM, AddU32_Slow) }
func TestSubU32_ASM_Quick(t *testing.T) { testOp(t, SubU32_ASM, SubU32_Slow) }
func TestMulU32_ASM_Quick(t *testing.T) { testOp(t, MulU32_ASM, MulU32_Slow) }

func benchpair(size int) ([]uint32, []uint32) {
	rng := rand.NewSource(0)
	dst := make([]uint32, size)
	src := make([]uint32, size)
	for i := range dst {
		dst[i] = uint32(rng.Int63())
		src[i] = uint32(rng.Int63())
	}
	return dst, src
}

func bench(b *testing.B, op op) {
	dst, src := benchpair(1 << 20)
	b.SetBytes(4 << 20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		op(dst, src)
	}
}

func BenchmarkAddU32_Slow(b *testing.B) { bench(b, AddU32_Slow) }
func BenchmarkAddU32_ASM(b *testing.B)  { bench(b, AddU32_ASM) }

func BenchmarkSubU32_Slow(b *testing.B) { bench(b, SubU32_Slow) }
func BenchmarkSubU32_ASM(b *testing.B)  { bench(b, SubU32_ASM) }

func BenchmarkMulU32_Slow(b *testing.B) { bench(b, MulU32_Slow) }
func BenchmarkMulU32_ASM(b *testing.B)  { bench(b, MulU32_ASM) }
