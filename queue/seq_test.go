package queue

import "testing"

func TestSeqSpinning(t *testing.T) {
	test(t, func(size int) Queue { return NewSeqSpinning(size) })
}

func BenchmarkSeqSpinning(b *testing.B) {
	bench(b, func(size int) Queue { return NewSeqSpinning(size) })
}

func TestSeqwSpinning(t *testing.T) {
	test(t, func(size int) Queue { return NewSeqwSpinning(size) })
}

func BenchmarkSeqwSpinning(b *testing.B) {
	bench(b, func(size int) Queue { return NewSeqwSpinning(size) })
}
