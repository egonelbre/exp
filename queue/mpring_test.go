package queue

import (
	"strconv"
	"testing"
)

func TestMPRing(t *testing.T) {
	t.Skip("broken")
	for _, batchSize := range BatchSizes {
		t.Run(strconv.Itoa(batchSize), func(t *testing.T) {
			test(t, func(size int) Queue { return NewMPRing(batchSize, size) })
		})
	}
}

func BenchmarkMPRing(b *testing.B) {
	b.Skip("broken")
	for _, batchSize := range BatchSizes {
		b.Run(strconv.Itoa(batchSize), func(b *testing.B) {
			bench(b, func(size int) Queue { return NewMPRing(batchSize, size) })
		})
	}
}

/*
func TestMPRingSpinning(t *testing.T) {
	for _, batchSize := range BatchSizes {
		t.Run(strconv.Itoa(batchSize), func(t *testing.T) {
			test(t, func(size int) Queue { return NewMPRingSpinning(batchSize, size) })
		})
	}
}

func BenchmarkMPRingSpinning(b *testing.B) {
	for _, batchSize := range BatchSizes {
		b.Run(strconv.Itoa(batchSize), func(b *testing.B) {
			bench(b, func(size int) Queue { return NewMPRingSpinning(batchSize, size) })
		})
	}

}
*/
