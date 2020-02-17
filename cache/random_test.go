package cache_test

import (
	"encoding/binary"
	"testing"

	"github.com/egonelbre/exp/cache"
	"github.com/loov/hrtime"
)

func BenchmarkRandom_Growing(b *testing.B) {
	b.ResetTimer()
	start := hrtime.Now()
	for k := 0; k < b.N; k++ {
		c := cache.NewRandom(1e9, nil)
		for i := 0; i < ItemsToAdd; i++ {
			var key [8]byte
			binary.LittleEndian.PutUint64(key[:], uint64(i))
			c.Add(cache.Hash(key[:]), 1)
		}
	}
	finish := hrtime.Now()
	b.StopTimer()
	b.ReportMetric(float64((finish-start).Nanoseconds())/float64(b.N*ItemsToAdd), "ns/item")
}

func BenchmarkRandom_Prealloc(b *testing.B) {
	b.ResetTimer()
	start := hrtime.Now()
	for k := 0; k < b.N; k++ {
		c := cache.NewRandomPrealloc(ItemsToAdd, 1e9, nil)
		for i := 0; i < ItemsToAdd; i++ {
			var key [8]byte
			binary.LittleEndian.PutUint64(key[:], uint64(i))
			c.Add(cache.Hash(key[:]), 1)
		}
	}
	finish := hrtime.Now()
	b.StopTimer()
	b.ReportMetric(float64((finish-start).Nanoseconds())/float64(b.N*ItemsToAdd), "ns/item")
}
