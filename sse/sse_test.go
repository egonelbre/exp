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

func simpleTest(t *testing.T, fast, slow op) {
	dst := []uint32{5, 4, 3, 2, 1, 8, 2, 8, 3, 3, 2, 5, 8, 5, 1, 8, 2, 8, 3, 3, 2, 5, 8}
	src := []uint32{1, 2, 5, 7, 5, 1, 4, 1, 5, 0, 5, 1, 1, 5, 1, 4, 1, 5, 0, 5, 1}

	dst_fast := make([]uint32, len(dst))
	dst_slow := make([]uint32, len(dst))
	copy(dst_fast, dst)
	copy(dst_slow, dst)

	fast(dst_fast, src)
	slow(dst_slow, src)

	if !same(dst_slow, dst_fast) {
		t.Errorf("failed\ndst\t%v\nsrc\t%v\nexp\t%v\ngot\t%v", dst, src, dst_slow, dst_fast)
	}
}

func TestAddU32_ASM_Simple(t *testing.T) { simpleTest(t, AddU32_ASM, AddU32_Slow) }
func TestSubU32_ASM_Simple(t *testing.T) { simpleTest(t, SubU32_ASM, SubU32_Slow) }
func TestMulU32_ASM_Simple(t *testing.T) { simpleTest(t, MulU32_ASM, MulU32_Slow) }

func TestAddU32_ASM_Quick(t *testing.T) { testOp(t, AddU32_ASM, AddU32_Slow) }
func TestSubU32_ASM_Quick(t *testing.T) { testOp(t, SubU32_ASM, SubU32_Slow) }
func TestMulU32_ASM_Quick(t *testing.T) { testOp(t, MulU32_ASM, MulU32_Slow) }

func TestAddU32_SSE_Simple(t *testing.T) { simpleTest(t, AddU32_SSE, AddU32_Slow) }
func TestSubU32_SSE_Simple(t *testing.T) { simpleTest(t, SubU32_SSE, SubU32_Slow) }
func TestMulU32_SSE_Simple(t *testing.T) { simpleTest(t, MulU32_SSE, MulU32_Slow) }

func TestAddU32_SSE_Quick(t *testing.T) { testOp(t, AddU32_SSE, AddU32_Slow) }
func TestSubU32_SSE_Quick(t *testing.T) { testOp(t, SubU32_SSE, SubU32_Slow) }
func TestMulU32_SSE_Quick(t *testing.T) { testOp(t, MulU32_SSE, MulU32_Slow) }

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
func BenchmarkAddU32_SSE(b *testing.B)  { bench(b, AddU32_SSE) }

func BenchmarkSubU32_Slow(b *testing.B) { bench(b, SubU32_Slow) }
func BenchmarkSubU32_ASM(b *testing.B)  { bench(b, SubU32_ASM) }
func BenchmarkSubU32_SSE(b *testing.B)  { bench(b, SubU32_SSE) }

func BenchmarkMulU32_Slow(b *testing.B) { bench(b, MulU32_Slow) }
func BenchmarkMulU32_ASM(b *testing.B)  { bench(b, MulU32_ASM) }
func BenchmarkMulU32_SSE(b *testing.B)  { bench(b, MulU32_SSE) }
