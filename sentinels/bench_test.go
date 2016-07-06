package sentinels

import "testing"

var (
	Items   = make([]int, 1<<20)
	Element = 1534
)

func BenchmarkSearch_Basic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchBasic(Items, Element)
	}
}

func BenchmarkSearch_Break(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchBreak(Items, Element)
	}
}

func BenchmarkSearch_Sentinel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchSentinel(Items, Element)
	}
}

func BenchmarkSearch_Unsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchUnsafe(Items, Element)
	}
}

func BenchmarkPartition_Break(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartitionBreak(Items, i%len(Items))
	}
}

func BenchmarkPartition_Sentinel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartitionSentinel(Items, i%len(Items))
	}
}
