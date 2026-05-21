package mul64

import (
	"math/bits"
	"testing"
)

// Sinks prevent dead-code elimination.
var (
	SinkHi, SinkLo uint64
)

// =============================================================================
// Correctness: confirm every shift variant matches bits.Mul64.
// Run with `go test ./...` before trusting benchmark results.
// =============================================================================

func TestVariantsAgreeWithMul64(t *testing.T) {
	inputs := []uint64{
		0, 1, 2, 3, 1<<32 - 1, 1 << 32, 1<<32 + 1,
		0xDEADBEEFCAFEBABE, 0x123456789ABCDEF0,
		0xFFFFFFFFFFFFFFFF, 0x8000000000000000, 0x7FFFFFFFFFFFFFFF,
	}
	for _, v := range variants {
		t.Run(v.name, func(t *testing.T) {
			for _, x := range inputs {
				wantHi, wantLo := v.mul(x)
				gotHi, gotLo := v.sft(x)
				if gotHi != wantHi || gotLo != wantLo {
					t.Errorf("x=%#016x: mul=(%#016x,%#016x) sft=(%#016x,%#016x)",
						x, wantHi, wantLo, gotHi, gotLo)
				}
			}
		})
	}
}

// =============================================================================
// Latency benchmark: each iteration depends on the previous result.
// This measures the LATENCY of one mul/shift sequence, i.e. how long the CPU
// takes when the next instruction must wait on this one. Relevant for code
// where the (hi, lo) feed immediately into a carry chain that feeds the next
// mul, which is the case in serial Montgomery-form scalar arithmetic.
// =============================================================================

func benchLatency(b *testing.B, f func(uint64) (uint64, uint64)) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		h, l := f(x)
		x = h ^ l ^ uint64(i) // forces dependency chain
	}
	SinkHi, SinkLo = f(x)
}

func BenchmarkLatency(b *testing.B) {
	for _, v := range variants {
		b.Run(v.name+"/Mul", func(b *testing.B) { benchLatency(b, v.mul) })
		b.Run(v.name+"/Sft", func(b *testing.B) { benchLatency(b, v.sft) })
	}
}

// =============================================================================
// Throughput benchmark: 8 independent calls per iteration with no carried
// dependency. This measures how many mul/shift sequences the CPU can sustain
// per unit time, which is relevant for code where many independent
// multiplications happen back-to-back (e.g. limb-by-limb field reduction).
// =============================================================================

var benchInputsX8 = [8]uint64{
	0xDEADBEEFCAFEBABE, 0x1234567890ABCDEF,
	0xFEDCBA9876543210, 0x0123456789ABCDEF,
	0xA1B2C3D4E5F60718, 0x9988776655443322,
	0x1122334455667788, 0xFFEEDDCCBBAA9988,
}

//go:noinline
func runX8(xs [8]uint64, f func(uint64) (uint64, uint64)) uint64 {
	var sum uint64
	for i := 0; i < 8; i++ {
		h, l := f(xs[i])
		sum ^= h ^ l
	}
	return sum
}

func benchThroughput(b *testing.B, f func(uint64) (uint64, uint64)) {
	xs := benchInputsX8
	for i := 0; i < b.N; i++ {
		SinkLo = runX8(xs, f)
	}
}

func BenchmarkThroughput(b *testing.B) {
	for _, v := range variants {
		b.Run(v.name+"/Mul", func(b *testing.B) { benchThroughput(b, v.mul) })
		b.Run(v.name+"/Sft", func(b *testing.B) { benchThroughput(b, v.sft) })
	}
}

// =============================================================================
// Reduction-style benchmark: 4 sequential dependent calls per iteration,
// matching the shape of one Montgomery reduction round (mu = lo * mu0, then
// four mul-and-accumulate steps). Closer to real crypto code than either of
// the above.
// =============================================================================

//go:noinline
func reduction(x uint64, f func(uint64) (uint64, uint64)) uint64 {
	h, l := f(x)
	for i := 0; i < 3; i++ {
		// Mix hi/lo back into x so each call depends on the previous.
		h2, l2 := f(h ^ l)
		// fold a Add-with-carry to look more like a real reduction step
		_, c := bits.Add64(l, h2, 0)
		h, l = h2, l2^c
	}
	return h ^ l
}

func benchReduction(b *testing.B, f func(uint64) (uint64, uint64)) {
	x := uint64(0xDEADBEEFCAFEBABE)
	for i := 0; i < b.N; i++ {
		x = reduction(x^uint64(i), f)
	}
	SinkLo = x
}

func BenchmarkReduction(b *testing.B) {
	for _, v := range variants {
		b.Run(v.name+"/Mul", func(b *testing.B) { benchReduction(b, v.mul) })
		b.Run(v.name+"/Sft", func(b *testing.B) { benchReduction(b, v.sft) })
	}
}
