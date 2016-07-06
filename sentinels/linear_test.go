package sentinels

import "testing"

var (
	Items   = make([]int, 1<<20)
	Element = 1534
)

func BenchmarkSearchBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchBasic(Items, Element)
	}
}

func BenchmarkSearchBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchBreak(Items, Element)
	}
}

func BenchmarkSearchSentinel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchSentinel(Items, Element)
	}
}

func BenchmarkSearchUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchUnsafe(Items, Element)
	}
}
