package mul64

import (
	"math/bits"
	"testing"
)

// =============================================================================
// "Raw" variants: a true MUL+UMULH baseline.
//
// Important finding: the Go compiler ALREADY does partial strength reduction
// for bits.Mul64(x, K) when K is a compile-time constant — for example
// bits.Mul64(x, 0xFFFFFFFFFFFFFFFF) compiles to `NEG` (for lo) + `UMULH x, -1`
// (for hi), not a real MUL+UMULH pair. So mul<Name> is comparing partial-SR
// against full-SR (the sft<Name> variants), not MUL vs SR.
//
// To benchmark pure MUL+UMULH, we route the constant through a runtime
// parameter so the compiler can't constant-fold the multiply. This gives a
// three-way comparison:
//
//   raw<Name>: full MUL+UMULH (constant hidden behind a parameter).
//   mul<Name>: bits.Mul64 with the constant inline (compiler does what it can).
//   sft<Name>: explicit full shift/sub strength reduction.
//
// On any given CPU, the answer to "is strength reduction worth it" really has
// two sub-questions:
//   (a) raw vs mul:  does the compiler's existing partial SR help?
//   (b) mul vs sft:  does additional, full SR (no multiplier at all) help?
// (b) is the relevant question for the SSA strength-reduction work — that's
// what an SSA-level rewrite would add on top of (a).
// =============================================================================

//go:noinline
func mulRaw(x, k uint64) (hi, lo uint64) { return bits.Mul64(x, k) }

func BenchmarkRaw_AllOnes(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := mulRaw(x, 0xFFFFFFFFFFFFFFFF)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkRaw_NegPow32(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := mulRaw(x, 0xFFFFFFFF00000000)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkRaw_NegTwo(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := mulRaw(x, 0xFFFFFFFFFFFFFFFE)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkRaw_Pow32(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := mulRaw(x, 1<<32)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkRaw_PlusOne(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := mulRaw(x, 0x100000001)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkRaw_Pow32Minus1(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := mulRaw(x, 0xFFFFFFFF)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

// =============================================================================
// Inline raw — defeat constant-folding by reading k from a sink the compiler
// can't reason about (a package-level variable assigned from outside the
// benchmark function). This gives an inline (no CALL+RET) MUL+UMULH baseline.
// =============================================================================

// RawK is set by an init-style helper to defeat constant tracking. Marked
// var (not const) and assigned in a //go:noinline helper so the optimizer
// can't see the value at the call site.
var RawK uint64

//go:noinline
func setRawK(k uint64) { RawK = k }

func BenchmarkInline_Raw_AllOnes(b *testing.B) {
	setRawK(0xFFFFFFFFFFFFFFFF)
	x := uint64(0xDEADBEEFCAFEBABE)
	k := RawK
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, k)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Raw_NegPow32(b *testing.B) {
	setRawK(0xFFFFFFFF00000000)
	x := uint64(0xDEADBEEFCAFEBABE)
	k := RawK
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, k)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Raw_Pow32(b *testing.B) {
	setRawK(1 << 32)
	x := uint64(0xDEADBEEFCAFEBABE)
	k := RawK
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, k)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}
