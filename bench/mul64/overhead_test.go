package mul64

import (
	"math/bits"
	"testing"
)

// =============================================================================
// Calibration: measure the cost of the bench harness itself so we can subtract
// it from the variant timings.
//
//   - identity:      direct, no //go:noinline (will be inlined)
//   - identityNoIn:  direct, //go:noinline (CALL + RET)
//   - identityVia:   indirect (function-pointer dispatch, what the variants use)
//
// The difference between identityVia and identityNoIn isolates the indirect-
// call overhead. The difference between identityNoIn and identity isolates the
// CALL/RET cost. Subtract these from the variant timings to estimate the true
// per-op cost of the mul/shift sequence.
// =============================================================================

func identity(x uint64) (hi, lo uint64) { return 0, x }

//go:noinline
func identityNoIn(x uint64) (hi, lo uint64) { return 0, x }

func BenchmarkOverhead_Identity_Inline(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := identity(x)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkOverhead_Identity_NoInline(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := identityNoIn(x)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkOverhead_Identity_Indirect(b *testing.B) {
	// Indirect call through a func value — same shape as the variant table.
	f := func(x uint64) (uint64, uint64) { return 0, x }
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := f(x)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

// =============================================================================
// Inline-friendly variants: same operations as in mul64.go, but written so the
// Go compiler can inline them into the bench loop. No //go:noinline, no
// callback. This gives the *true* per-op cost without harness overhead.
//
// Compare these timings to the Latency benchmarks in mul64_test.go to see how
// much call overhead distorted the measurement.
// =============================================================================

func BenchmarkInline_Latency_AllOnes_Mul(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, 0xFFFFFFFFFFFFFFFF)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_AllOnes_Sft(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		lo, br := bits.Sub64(0, x, 0)
		hi := x - br
		x = hi ^ lo ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_NegPow32_Mul(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, 0xFFFFFFFF00000000)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_NegPow32_Sft(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		lo, br := bits.Sub64(0, x<<32, 0)
		hi := x - (x >> 32) - br
		x = hi ^ lo ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_Pow32_Mul(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, 1<<32)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_Pow32_Sft(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		hi := x >> 32
		lo := x << 32
		x = hi ^ lo ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_PlusOne_Mul(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := bits.Mul64(x, 0x100000001)
		x = h ^ l ^ uint64(i)
	}
	SinkLo = x
}

func BenchmarkInline_Latency_PlusOne_Sft(b *testing.B) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		lo, c := bits.Add64(x<<32, x, 0)
		hi := (x >> 32) + c
		x = hi ^ lo ^ uint64(i)
	}
	SinkLo = x
}
