package mul64

import (
	"math/bits"
	"testing"
)

// =============================================================================
// Kernel benchmarks: measure the strength reduction in the context of
// surrounding work, not in isolation.
//
// The isolated micro-benchmarks (BenchmarkLatency / BenchmarkThroughput /
// BenchmarkInline_*) measure one operation at a time. Real crypto code mixes
// many operations per kernel, and the bottleneck depends on which execution
// units saturate first:
//
//   - If the multiplier units are the bottleneck, replacing one MUL with shifts
//     frees a multiplier to issue another MUL in parallel → speedup.
//   - If the ALU / carry chain is the bottleneck, replacing one MUL with 3-5
//     ALU ops adds pressure → slowdown.
//
// These kernels approximate one round of CIOS Montgomery reduction by a
// 4-limb modulus where two of the four limbs are sparse (the typical NIST
// shape: P-256 ord, P-384 field, etc.). Each kernel does the same set of
// "other" operations; the only difference between the Mul and Sft variants
// is whether the two sparse-constant multiplications use bits.Mul64 or the
// strength-reduced shift sequence.
// =============================================================================

// MU is a fake "Montgomery factor" mod ord — any dense 64-bit constant works
// for showing that the dense mul doesn't get strength-reduced.
const MU = 0xCCD1C8AAEE00BC4F

// =============================================================================
// reductionRound: one round of word-by-word CIOS reduction by a 4-limb modulus
// (limb 0 = dense, limb 1 = dense, limb 2 = 0xFFFFFFFFFFFFFFFF, limb 3 =
// 0xFFFFFFFF00000000). The mu-by-modulus pairs in limbs 2 and 3 are what
// strength reduction targets. Inputs t0..t3 are the four high accumulator
// limbs from a prior step.
// =============================================================================

//go:noinline
func reductionRoundMul(t0, t1, t2, t3, lo uint64) (r0, r1, r2, r3, r4 uint64) {
	// mu = lo * MU mod 2^64  (dense, no SR)
	mu, _ := bits.Mul64(lo, MU)
	_ = mu // suppress hi
	mu = lo * MU

	// mu * m[0]  (dense, not strength-reducible)
	h0, l0 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
	// mu * m[1]  (dense, not strength-reducible)
	h1, l1 := bits.Mul64(mu, 0xBCE6FAADA7179E84)
	// mu * m[2] = mu * (2^64 - 1)   <- sparse, SR target
	h2, l2 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
	// mu * m[3] = mu * (2^64 - 2^32) <- sparse, SR target
	h3, l3 := bits.Mul64(mu, 0xFFFFFFFF00000000)

	// Carry chain accumulating into (r0..r4).
	var c uint64
	r0, c = bits.Add64(t0, l0, 0)
	r1, c = bits.Add64(t1, l1, c)
	r2, c = bits.Add64(t2, l2, c)
	r3, c = bits.Add64(t3, l3, c)
	r4 = c

	// And the high halves, shifted by one limb.
	r1, c = bits.Add64(r1, h0, 0)
	r2, c = bits.Add64(r2, h1, c)
	r3, c = bits.Add64(r3, h2, c)
	r4, _ = bits.Add64(r4, h3, c)
	return
}

//go:noinline
func reductionRoundSft(t0, t1, t2, t3, lo uint64) (r0, r1, r2, r3, r4 uint64) {
	mu := lo * MU

	// mu * m[0]  (dense, unchanged)
	h0, l0 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
	// mu * m[1]  (dense, unchanged)
	h1, l1 := bits.Mul64(mu, 0xBCE6FAADA7179E84)
	// mu * (2^64 - 1) via strength reduction
	l2, br := bits.Sub64(0, mu, 0)
	h2 := mu - br
	// mu * (2^64 - 2^32) via strength reduction
	l3, br2 := bits.Sub64(0, mu<<32, 0)
	h3 := mu - (mu >> 32) - br2

	var c uint64
	r0, c = bits.Add64(t0, l0, 0)
	r1, c = bits.Add64(t1, l1, c)
	r2, c = bits.Add64(t2, l2, c)
	r3, c = bits.Add64(t3, l3, c)
	r4 = c

	r1, c = bits.Add64(r1, h0, 0)
	r2, c = bits.Add64(r2, h1, c)
	r3, c = bits.Add64(r3, h2, c)
	r4, _ = bits.Add64(r4, h3, c)
	return
}

func BenchmarkKernel_ReductionRound_Mul(b *testing.B) {
	t0, t1, t2, t3, lo := uint64(1), uint64(2), uint64(3), uint64(4), uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		var r4 uint64
		t0, t1, t2, t3, r4 = reductionRoundMul(t0, t1, t2, t3, lo)
		lo = r4 ^ uint64(i) // feed back to defeat OoO speculation across iterations
	}
	SinkLo = t0 ^ t1 ^ t2 ^ t3 ^ lo
}

func BenchmarkKernel_ReductionRound_Sft(b *testing.B) {
	t0, t1, t2, t3, lo := uint64(1), uint64(2), uint64(3), uint64(4), uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		var r4 uint64
		t0, t1, t2, t3, r4 = reductionRoundSft(t0, t1, t2, t3, lo)
		lo = r4 ^ uint64(i)
	}
	SinkLo = t0 ^ t1 ^ t2 ^ t3 ^ lo
}

// =============================================================================
// Independent-rounds kernel: 4 reduction rounds with NO data dependency between
// them. This is the throughput-bound version — it exposes how much
// parallelism the CPU can extract when many sparse muls happen back-to-back.
// In real Montgomery code rounds are sequential, but this isolates the
// "many independent sparse muls" case (which is what hurt P-384/P-521).
// =============================================================================

//go:noinline
func fourRoundsMul(in [5]uint64) uint64 {
	t0, t1, t2, t3, t4 := reductionRoundMul(in[0], in[1], in[2], in[3], in[4])
	u0, u1, u2, u3, u4 := reductionRoundMul(in[1], in[2], in[3], in[4], in[0])
	v0, v1, v2, v3, v4 := reductionRoundMul(in[2], in[3], in[4], in[0], in[1])
	w0, w1, w2, w3, w4 := reductionRoundMul(in[3], in[4], in[0], in[1], in[2])
	return t0 ^ t1 ^ t2 ^ t3 ^ t4 ^ u0 ^ u1 ^ u2 ^ u3 ^ u4 ^ v0 ^ v1 ^ v2 ^ v3 ^ v4 ^ w0 ^ w1 ^ w2 ^ w3 ^ w4
}

//go:noinline
func fourRoundsSft(in [5]uint64) uint64 {
	t0, t1, t2, t3, t4 := reductionRoundSft(in[0], in[1], in[2], in[3], in[4])
	u0, u1, u2, u3, u4 := reductionRoundSft(in[1], in[2], in[3], in[4], in[0])
	v0, v1, v2, v3, v4 := reductionRoundSft(in[2], in[3], in[4], in[0], in[1])
	w0, w1, w2, w3, w4 := reductionRoundSft(in[3], in[4], in[0], in[1], in[2])
	return t0 ^ t1 ^ t2 ^ t3 ^ t4 ^ u0 ^ u1 ^ u2 ^ u3 ^ u4 ^ v0 ^ v1 ^ v2 ^ v3 ^ v4 ^ w0 ^ w1 ^ w2 ^ w3 ^ w4
}

func BenchmarkKernel_FourRoundsParallel_Mul(b *testing.B) {
	in := [5]uint64{0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0x9ABCDEF0, 0xFEDCBA98}
	for i := 0; i < b.N; i++ {
		SinkLo = fourRoundsMul(in)
		in[0] ^= uint64(i)
	}
}

func BenchmarkKernel_FourRoundsParallel_Sft(b *testing.B) {
	in := [5]uint64{0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0x9ABCDEF0, 0xFEDCBA98}
	for i := 0; i < b.N; i++ {
		SinkLo = fourRoundsSft(in)
		in[0] ^= uint64(i)
	}
}

// =============================================================================
// Inline kernel: same shape as reductionRound, but the WHOLE round is in the
// benchmark loop body — no function-call boundary. Strips the CALL+RET cost
// from the comparison and matches how an inliner-friendly compiler would lay
// the code out.
// =============================================================================

func BenchmarkKernel_Inline_ReductionRound_Mul(b *testing.B) {
	t0, t1, t2, t3, lo := uint64(1), uint64(2), uint64(3), uint64(4), uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		mu := lo * MU
		h0, l0 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
		h1, l1 := bits.Mul64(mu, 0xBCE6FAADA7179E84)
		h2, l2 := bits.Mul64(mu, 0xFFFFFFFFFFFFFFFF)
		h3, l3 := bits.Mul64(mu, 0xFFFFFFFF00000000)
		var c, r4 uint64
		t0, c = bits.Add64(t0, l0, 0)
		t1, c = bits.Add64(t1, l1, c)
		t2, c = bits.Add64(t2, l2, c)
		t3, c = bits.Add64(t3, l3, c)
		r4 = c
		t1, c = bits.Add64(t1, h0, 0)
		t2, c = bits.Add64(t2, h1, c)
		t3, c = bits.Add64(t3, h2, c)
		r4, _ = bits.Add64(r4, h3, c)
		lo = r4 ^ uint64(i)
	}
	SinkLo = t0 ^ t1 ^ t2 ^ t3 ^ lo
}

func BenchmarkKernel_Inline_ReductionRound_Sft(b *testing.B) {
	t0, t1, t2, t3, lo := uint64(1), uint64(2), uint64(3), uint64(4), uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		mu := lo * MU
		h0, l0 := bits.Mul64(mu, 0xF3B9CAC2FC632551)
		h1, l1 := bits.Mul64(mu, 0xBCE6FAADA7179E84)
		l2, br := bits.Sub64(0, mu, 0)
		h2 := mu - br
		l3, br2 := bits.Sub64(0, mu<<32, 0)
		h3 := mu - (mu >> 32) - br2
		var c, r4 uint64
		t0, c = bits.Add64(t0, l0, 0)
		t1, c = bits.Add64(t1, l1, c)
		t2, c = bits.Add64(t2, l2, c)
		t3, c = bits.Add64(t3, l3, c)
		r4 = c
		t1, c = bits.Add64(t1, h0, 0)
		t2, c = bits.Add64(t2, h1, c)
		t3, c = bits.Add64(t3, h2, c)
		r4, _ = bits.Add64(r4, h3, c)
		lo = r4 ^ uint64(i)
	}
	SinkLo = t0 ^ t1 ^ t2 ^ t3 ^ lo
}

// =============================================================================
// Correctness: Mul and Sft variants of each kernel must produce identical
// results for arbitrary inputs.
// =============================================================================

func TestKernelEquivalence(t *testing.T) {
	inputs := [][5]uint64{
		{0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5},
		{0xDEADBEEF, 0xCAFEBABE, 0x12345678, 0x9ABCDEF0, 0xFEDCBA98},
		{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
		{0x8000000000000000, 0, 0, 0, 0x1},
	}
	for _, in := range inputs {
		m0, m1, m2, m3, m4 := reductionRoundMul(in[0], in[1], in[2], in[3], in[4])
		s0, s1, s2, s3, s4 := reductionRoundSft(in[0], in[1], in[2], in[3], in[4])
		if m0 != s0 || m1 != s1 || m2 != s2 || m3 != s3 || m4 != s4 {
			t.Errorf("input %v:\n  mul=(%x,%x,%x,%x,%x)\n  sft=(%x,%x,%x,%x,%x)",
				in, m0, m1, m2, m3, m4, s0, s1, s2, s3, s4)
		}
		if fourRoundsMul(in) != fourRoundsSft(in) {
			t.Errorf("fourRounds mismatch on %v", in)
		}
	}
}
