package mul64

import (
	"math/bits"
	"testing"
)

// =============================================================================
// 6-limb reduction round (P-384 shape).
//
// P-384's modulus has limbs: 0xFFFFFFFFFFFFFFFF, 0x00000000FFFFFFFF (= 2^32-1),
// 0xFFFFFFFFFFFFFFFE (= -2), 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF.
// All six limbs are sparse — opposite extreme from the 4-limb P-256 ord
// modulus which is 2-sparse-out-of-4. This is where the SSA-level rewrite
// regressed on real code; this kernel directly models that case.
// =============================================================================

//go:noinline
func reductionRound6Mul(t [6]uint64, lo uint64) (r [7]uint64) {
	mu := lo * MU
	h0, l0 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h1, l1 := bits.Mul64(mu, 0x00000000FFFFFFFF)
	h2, l2 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFE)
	h3, l3 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h4, l4 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h5, l5 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)

	var c uint64
	r[0], c = bits.Add64(t[0], l0, 0)
	r[1], c = bits.Add64(t[1], l1, c)
	r[2], c = bits.Add64(t[2], l2, c)
	r[3], c = bits.Add64(t[3], l3, c)
	r[4], c = bits.Add64(t[4], l4, c)
	r[5], c = bits.Add64(t[5], l5, c)
	r[6] = c
	r[1], c = bits.Add64(r[1], h0, 0)
	r[2], c = bits.Add64(r[2], h1, c)
	r[3], c = bits.Add64(r[3], h2, c)
	r[4], c = bits.Add64(r[4], h3, c)
	r[5], c = bits.Add64(r[5], h4, c)
	r[6], _ = bits.Add64(r[6], h5, c)
	return
}

//go:noinline
func reductionRound6Sft(t [6]uint64, lo uint64) (r [7]uint64) {
	mu := lo * MU
	// 0xFFFFFFFFFFFFFFFF
	l0, br0 := bits.Sub64(0, mu, 0)
	h0 := mu - br0
	// 0x00000000FFFFFFFF = 2^32 - 1
	l1, br1 := bits.Sub64(mu<<32, mu, 0)
	h1 := (mu >> 32) - br1
	// 0xFFFFFFFFFFFFFFFE = -2 = 2^64 - 2^1
	l2, br2 := bits.Sub64(0, mu<<1, 0)
	h2 := mu - (mu >> 63) - br2
	// three more 0xFFFFFFFFFFFFFFFF
	l3, br3 := bits.Sub64(0, mu, 0)
	h3 := mu - br3
	l4, br4 := bits.Sub64(0, mu, 0)
	h4 := mu - br4
	l5, br5 := bits.Sub64(0, mu, 0)
	h5 := mu - br5

	var c uint64
	r[0], c = bits.Add64(t[0], l0, 0)
	r[1], c = bits.Add64(t[1], l1, c)
	r[2], c = bits.Add64(t[2], l2, c)
	r[3], c = bits.Add64(t[3], l3, c)
	r[4], c = bits.Add64(t[4], l4, c)
	r[5], c = bits.Add64(t[5], l5, c)
	r[6] = c
	r[1], c = bits.Add64(r[1], h0, 0)
	r[2], c = bits.Add64(r[2], h1, c)
	r[3], c = bits.Add64(r[3], h2, c)
	r[4], c = bits.Add64(r[4], h3, c)
	r[5], c = bits.Add64(r[5], h4, c)
	r[6], _ = bits.Add64(r[6], h5, c)
	return
}

func BenchmarkKernel_P384Round_Mul(b *testing.B) {
	t := [6]uint64{1, 2, 3, 4, 5, 6}
	lo := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		r := reductionRound6Mul(t, lo)
		t = [6]uint64{r[0], r[1], r[2], r[3], r[4], r[5]}
		lo = r[6] ^ uint64(i)
	}
	SinkLo = t[0] ^ t[1] ^ t[2] ^ t[3] ^ t[4] ^ t[5] ^ lo
}

func BenchmarkKernel_P384Round_Sft(b *testing.B) {
	t := [6]uint64{1, 2, 3, 4, 5, 6}
	lo := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		r := reductionRound6Sft(t, lo)
		t = [6]uint64{r[0], r[1], r[2], r[3], r[4], r[5]}
		lo = r[6] ^ uint64(i)
	}
	SinkLo = t[0] ^ t[1] ^ t[2] ^ t[3] ^ t[4] ^ t[5] ^ lo
}

// =============================================================================
// Sparse-density sweep: 4-limb reduction round with 1, 2, 3, or 4 of the four
// constant multiplications being strength-reducible. Shows the crossover —
// how many sparse limbs you need before SR becomes a win (or loss).
//
// Constants used (in order of "sparse slot fill"):
//   slot 0: 0xF3B9CAC2FC632551 (dense)
//   slot 1: 0xBCE6FAADA7179E84 (dense)
//   slot 2: 0xFFFFFFFFFFFFFFFF (sparse: 2^64-1)
//   slot 3: 0xFFFFFFFF00000000 (sparse: 2^64-2^32)
//
// Each round-Nx variant promotes N of the four slots to sparse constants.
// =============================================================================

// 1 sparse out of 4 (replace slot 0 with sparse)
//
//go:noinline
func reductionRound4_S1_Mul(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	h0, l0 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h1, l1 := bits.Mul64(mu, 0xBCE6FAADA7179E84)
	h2, l2 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
	h3, l3 := bits.Mul64(mu, 0xCCD1C8AAEE00BC4F)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

//go:noinline
func reductionRound4_S1_Sft(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	l0, br := bits.Sub64(0, mu, 0)
	h0 := mu - br
	h1, l1 := bits.Mul64(mu, 0xBCE6FAADA7179E84)
	h2, l2 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
	h3, l3 := bits.Mul64(mu, 0xCCD1C8AAEE00BC4F)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

// 2 sparse out of 4 — same as the original ReductionRound; reproduced here
// alongside the 1/3/4 variants so the sweep is one consistent table.
//
//go:noinline
func reductionRound4_S2_Mul(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	h0, l0 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h1, l1 := bits.Mul64(mu, 0xFFFFFFFF00000000)
	h2, l2 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
	h3, l3 := bits.Mul64(mu, 0xCCD1C8AAEE00BC4F)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

//go:noinline
func reductionRound4_S2_Sft(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	l0, br0 := bits.Sub64(0, mu, 0)
	h0 := mu - br0
	l1, br1 := bits.Sub64(0, mu<<32, 0)
	h1 := mu - (mu >> 32) - br1
	h2, l2 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
	h3, l3 := bits.Mul64(mu, 0xCCD1C8AAEE00BC4F)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

// 3 sparse out of 4
//
//go:noinline
func reductionRound4_S3_Mul(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	h0, l0 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h1, l1 := bits.Mul64(mu, 0xFFFFFFFF00000000)
	h2, l2 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFE)
	h3, l3 := bits.Mul64(mu, 0xCCD1C8AAEE00BC4F)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

//go:noinline
func reductionRound4_S3_Sft(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	l0, br0 := bits.Sub64(0, mu, 0)
	h0 := mu - br0
	l1, br1 := bits.Sub64(0, mu<<32, 0)
	h1 := mu - (mu >> 32) - br1
	l2, br2 := bits.Sub64(0, mu<<1, 0)
	h2 := mu - (mu >> 63) - br2
	h3, l3 := bits.Mul64(mu, 0xCCD1C8AAEE00BC4F)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

// 4 sparse out of 4
//
//go:noinline
func reductionRound4_S4_Mul(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	h0, l0 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	h1, l1 := bits.Mul64(mu, 0xFFFFFFFF00000000)
	h2, l2 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFE)
	h3, l3 := bits.Mul64(mu, 0x100000001)
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

//go:noinline
func reductionRound4_S4_Sft(t [4]uint64, lo uint64) (r [5]uint64) {
	mu := lo * MU
	l0, br0 := bits.Sub64(0, mu, 0)
	h0 := mu - br0
	l1, br1 := bits.Sub64(0, mu<<32, 0)
	h1 := mu - (mu >> 32) - br1
	l2, br2 := bits.Sub64(0, mu<<1, 0)
	h2 := mu - (mu >> 63) - br2
	l3, c3 := bits.Add64(mu<<32, mu, 0)
	h3 := (mu >> 32) + c3
	return reduceAccum4(t, l0, l1, l2, l3, h0, h1, h2, h3)
}

// reduceAccum4 is the shared carry-chain "tail" for the 4-limb variants —
// extracted so the body of each S_N variant focuses on the muls only.
func reduceAccum4(t [4]uint64, l0, l1, l2, l3, h0, h1, h2, h3 uint64) (r [5]uint64) {
	var c uint64
	r[0], c = bits.Add64(t[0], l0, 0)
	r[1], c = bits.Add64(t[1], l1, c)
	r[2], c = bits.Add64(t[2], l2, c)
	r[3], c = bits.Add64(t[3], l3, c)
	r[4] = c
	r[1], c = bits.Add64(r[1], h0, 0)
	r[2], c = bits.Add64(r[2], h1, c)
	r[3], c = bits.Add64(r[3], h2, c)
	r[4], _ = bits.Add64(r[4], h3, c)
	return
}

func benchDensityRound(b *testing.B, f func([4]uint64, uint64) [5]uint64) {
	t := [4]uint64{1, 2, 3, 4}
	lo := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		r := f(t, lo)
		t = [4]uint64{r[0], r[1], r[2], r[3]}
		lo = r[4] ^ uint64(i)
	}
	SinkLo = t[0] ^ t[1] ^ t[2] ^ t[3] ^ lo
}

func BenchmarkKernel_Density_S1_Mul(b *testing.B) {
	benchDensityRound(b, reductionRound4_S1_Mul)
}
func BenchmarkKernel_Density_S1_Sft(b *testing.B) {
	benchDensityRound(b, reductionRound4_S1_Sft)
}
func BenchmarkKernel_Density_S2_Mul(b *testing.B) {
	benchDensityRound(b, reductionRound4_S2_Mul)
}
func BenchmarkKernel_Density_S2_Sft(b *testing.B) {
	benchDensityRound(b, reductionRound4_S2_Sft)
}
func BenchmarkKernel_Density_S3_Mul(b *testing.B) {
	benchDensityRound(b, reductionRound4_S3_Mul)
}
func BenchmarkKernel_Density_S3_Sft(b *testing.B) {
	benchDensityRound(b, reductionRound4_S3_Sft)
}
func BenchmarkKernel_Density_S4_Mul(b *testing.B) {
	benchDensityRound(b, reductionRound4_S4_Mul)
}
func BenchmarkKernel_Density_S4_Sft(b *testing.B) {
	benchDensityRound(b, reductionRound4_S4_Sft)
}

// =============================================================================
// Chained Montgomery multiplication (4 rounds, sequential).
//
// Each round's accumulator output feeds the next round's input — the actual
// shape of `p256OrdMul`, where each iteration consumes one operand limb. This
// is the closest the bench comes to the real Montgomery mul, and shows how
// long sequential carry chains affect SR profitability.
// =============================================================================

//go:noinline
func chainedMontMulMul(operand [4]uint64, in [4]uint64) (out [4]uint64, top uint64) {
	t := in
	for k := 0; k < 4; k++ {
		// Outer step: accumulate operand[k] * x where x is some constant
		// (a single MUL per outer step). Here we use a fake "x" to keep the
		// kernel sane.
		hi, lo := bits.Mul64(operand[k], 0x123456789ABCDEF0)
		var c uint64
		t[0], c = bits.Add64(t[0], lo, 0)
		t[1], c = bits.Add64(t[1], hi, c)
		t[2], c = bits.Add64(t[2], 0, c)
		t[3], c = bits.Add64(t[3], 0, c)
		top += c

		// Reduction round on the accumulator.
		r := reductionRound4_S2_Mul(t, t[0])
		t = [4]uint64{r[1], r[2], r[3], r[4]}
		top ^= r[0] // fold the discarded low limb into top so it can't be DCEd
	}
	return t, top
}

//go:noinline
func chainedMontMulSft(operand [4]uint64, in [4]uint64) (out [4]uint64, top uint64) {
	t := in
	for k := 0; k < 4; k++ {
		hi, lo := bits.Mul64(operand[k], 0x123456789ABCDEF0)
		var c uint64
		t[0], c = bits.Add64(t[0], lo, 0)
		t[1], c = bits.Add64(t[1], hi, c)
		t[2], c = bits.Add64(t[2], 0, c)
		t[3], c = bits.Add64(t[3], 0, c)
		top += c

		r := reductionRound4_S2_Sft(t, t[0])
		t = [4]uint64{r[1], r[2], r[3], r[4]}
		top ^= r[0]
	}
	return t, top
}

func BenchmarkKernel_ChainedMontMul_Mul(b *testing.B) {
	operand := [4]uint64{0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0x9ABCDEF0}
	in := [4]uint64{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		out, top := chainedMontMulMul(operand, in)
		in = out
		in[0] ^= top ^ uint64(i)
	}
	SinkLo = in[0] ^ in[1] ^ in[2] ^ in[3]
}

func BenchmarkKernel_ChainedMontMul_Sft(b *testing.B) {
	operand := [4]uint64{0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0x9ABCDEF0}
	in := [4]uint64{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		out, top := chainedMontMulSft(operand, in)
		in = out
		in[0] ^= top ^ uint64(i)
	}
	SinkLo = in[0] ^ in[1] ^ in[2] ^ in[3]
}

// =============================================================================
// Correctness: every Mul / Sft pair must agree across the wider kernels.
// =============================================================================

func TestKernelMoreEquivalence(t *testing.T) {
	t.Run("P384Round", func(t *testing.T) {
		inputs := [][7]uint64{
			{0, 0, 0, 0, 0, 0, 0},
			{1, 2, 3, 4, 5, 6, 0xDEADBEEF},
			{0xFFFFFFFFFFFFFFFF, 0, 0xFFFFFFFFFFFFFFFF, 0, 0xFFFFFFFFFFFFFFFF, 0, 0xCAFEBABE},
		}
		for _, in := range inputs {
			tArr := [6]uint64{in[0], in[1], in[2], in[3], in[4], in[5]}
			lo := in[6]
			m := reductionRound6Mul(tArr, lo)
			s := reductionRound6Sft(tArr, lo)
			if m != s {
				t.Errorf("P384Round %v: mul=%v sft=%v", in, m, s)
			}
		}
	})
	t.Run("DensitySweep", func(t *testing.T) {
		variants := []struct {
			name     string
			mul, sft func([4]uint64, uint64) [5]uint64
		}{
			{"S1", reductionRound4_S1_Mul, reductionRound4_S1_Sft},
			{"S2", reductionRound4_S2_Mul, reductionRound4_S2_Sft},
			{"S3", reductionRound4_S3_Mul, reductionRound4_S3_Sft},
			{"S4", reductionRound4_S4_Mul, reductionRound4_S4_Sft},
		}
		inputs := [][5]uint64{
			{0, 0, 0, 0, 0},
			{1, 2, 3, 4, 0xDEADBEEF},
			{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xCAFEBABE},
		}
		for _, v := range variants {
			for _, in := range inputs {
				tArr := [4]uint64{in[0], in[1], in[2], in[3]}
				lo := in[4]
				m := v.mul(tArr, lo)
				s := v.sft(tArr, lo)
				if m != s {
					t.Errorf("%s input %v: mul=%v sft=%v", v.name, in, m, s)
				}
			}
		}
	})
	t.Run("ChainedMontMul", func(t *testing.T) {
		operand := [4]uint64{0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0x9ABCDEF0}
		inputs := [][4]uint64{
			{0, 0, 0, 0}, {1, 2, 3, 4},
			{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
		}
		for _, in := range inputs {
			mo, mt := chainedMontMulMul(operand, in)
			so, st := chainedMontMulSft(operand, in)
			if mo != so || mt != st {
				t.Errorf("ChainedMontMul %v: mul=(%v,%x) sft=(%v,%x)", in, mo, mt, so, st)
			}
		}
	})
}
