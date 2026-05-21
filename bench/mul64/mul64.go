// Package mul64 compares bits.Mul64(x, K) — a 64x64 -> 128 unsigned multiply
// — against equivalent shift/add/sub sequences ("strength reductions") for
// constants K with a sparse signed-digit form.
//
// The point is to measure, on whatever hardware you're on, whether replacing
// MUL+UMULH with a shift sequence is actually faster. This matters for
// compilers (e.g. the Go SSA backend) deciding whether to apply strength
// reduction for sparse moduli used in NIST/Solinas curve arithmetic.
//
// Each constant K has two variants:
//
//   - mul<Name>:   uses bits.Mul64 directly. Compiles to MUL+UMULH on arm64,
//                  MULQ on amd64.
//   - sft<Name>:   uses the explicit shift/sub sequence that strength reduction
//                  would produce.
//
// The variants are //go:noinline so the bench measures the body + call/return
// uniformly. Inlining is irrelevant for the question (both lower to the same
// number of instructions when inlined).
package mul64

import "math/bits"

// =============================================================================
// k = 2^64 - 1  (all ones; appears as the high limb of many NIST moduli).
// =============================================================================

//go:noinline
func mulAllOnes(x uint64) (hi, lo uint64) {
	return bits.Mul64(x, 0xFFFFFFFFFFFFFFFF)
}

//go:noinline
func sftAllOnes(x uint64) (hi, lo uint64) {
	lo, b := bits.Sub64(0, x, 0)
	hi = x - b
	return
}

// =============================================================================
// k = 2^64 - 2^32  (= -2^32). P-256 ord and P-384 field both use this.
// =============================================================================

//go:noinline
func mulNegPow32(x uint64) (hi, lo uint64) {
	return bits.Mul64(x, 0xFFFFFFFF00000000)
}

//go:noinline
func sftNegPow32(x uint64) (hi, lo uint64) {
	lo, b := bits.Sub64(0, x<<32, 0)
	hi = x - (x >> 32) - b
	return
}

// =============================================================================
// k = 2^64 - 2^1  (= -2). P-384 field uses this.
// =============================================================================

//go:noinline
func mulNegTwo(x uint64) (hi, lo uint64) {
	return bits.Mul64(x, 0xFFFFFFFFFFFFFFFE)
}

//go:noinline
func sftNegTwo(x uint64) (hi, lo uint64) {
	lo, b := bits.Sub64(0, x<<1, 0)
	hi = x - (x >> 63) - b
	return
}

// =============================================================================
// k = 2^32  (pure power of two). 2 shifts vs MUL+UMULH.
// =============================================================================

//go:noinline
func mulPow32(x uint64) (hi, lo uint64) {
	return bits.Mul64(x, 1<<32)
}

//go:noinline
func sftPow32(x uint64) (hi, lo uint64) {
	return x >> 32, x << 32
}

// =============================================================================
// k = 2^32 + 1  (two-term, positive). P-384 field uses this.
// =============================================================================

//go:noinline
func mulPlusOne(x uint64) (hi, lo uint64) {
	return bits.Mul64(x, 0x100000001)
}

//go:noinline
func sftPlusOne(x uint64) (hi, lo uint64) {
	lo, c := bits.Add64(x<<32, x, 0)
	hi = (x >> 32) + c
	return
}

// =============================================================================
// k = 2^32 - 1  (two-term, all-ones half-word).
// =============================================================================

//go:noinline
func mulPow32Minus1(x uint64) (hi, lo uint64) {
	return bits.Mul64(x, 0xFFFFFFFF)
}

//go:noinline
func sftPow32Minus1(x uint64) (hi, lo uint64) {
	lo, b := bits.Sub64(x<<32, x, 0)
	hi = (x >> 32) - b
	return
}

// =============================================================================
// Variant table: drives the benchmark loop. Add new patterns here.
// =============================================================================

type variant struct {
	name string
	mul  func(uint64) (hi, lo uint64)
	sft  func(uint64) (hi, lo uint64)
}

var variants = []variant{
	{"AllOnes_x_2_64-1", mulAllOnes, sftAllOnes},
	{"NegPow32_x_2_64-2_32", mulNegPow32, sftNegPow32},
	{"NegTwo_x_2_64-2", mulNegTwo, sftNegTwo},
	{"Pow32_x_2_32", mulPow32, sftPow32},
	{"PlusOne_x_2_32+1", mulPlusOne, sftPlusOne},
	{"Pow32Minus1_x_2_32-1", mulPow32Minus1, sftPow32Minus1},
}
