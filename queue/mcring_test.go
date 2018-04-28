package queue

import (
	"strconv"
	"testing"
)

func TestMCRing(t *testing.T) {
	for _, batchSize := range BatchSizes {
		t.Run(strconv.Itoa(batchSize), func(t *testing.T) {
			test(t, func(size int) Queue { return NewMCRing(batchSize, size) })
		})
	}
}

func BenchmarkMCRing(b *testing.B) {
	for _, batchSize := range BatchSizes {
		b.Run(strconv.Itoa(batchSize), func(b *testing.B) {
			bench(b, func(size int) Queue { return NewMCRing(batchSize, size) })
		})
	}

}

func TestMCRingSpinning(t *testing.T) {
	for _, batchSize := range BatchSizes {
		t.Run(strconv.Itoa(batchSize), func(t *testing.T) {
			test(t, func(size int) Queue { return NewMCRingSpinning(batchSize, size) })
		})
	}
}

func BenchmarkMCRingSpinning(b *testing.B) {
	for _, batchSize := range BatchSizes {
		b.Run(strconv.Itoa(batchSize), func(b *testing.B) {
			bench(b, func(size int) Queue { return NewMCRingSpinning(batchSize, size) })
		})
	}

}
