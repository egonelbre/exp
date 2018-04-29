package queue

import "testing"

func TestSeqSpinning(t *testing.T) {
	test(t, func(size int) Queue { return NewSeqSpinning(size) })
}

func BenchmarkSeqSpinning(b *testing.B) {
	bench(b, func(size int) Queue { return NewSeqSpinning(size) })
}

func TestSeqSPMCSpinning(t *testing.T) {
	test(t, func(size int) Queue { return NewSeqSPMCSpinning(size) })
}

func BenchmarkSeqSPMCSpinning(b *testing.B) {
	bench(b, func(size int) Queue { return NewSeqSPMCSpinning(size) })
}

func TestSeqwSpinning(t *testing.T) {
	test(t, func(size int) Queue { return NewSeqwSpinning(size) })
}

func BenchmarkSeqwSpinning(b *testing.B) {
	bench(b, func(size int) Queue { return NewSeqwSpinning(size) })
}
