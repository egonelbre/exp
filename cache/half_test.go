package cache_test

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"testing"

	"github.com/egonelbre/exp/cache"
	"github.com/loov/hrtime"
)

func ExampleHalf() {
	c := cache.NewHalf(1024, func(x cache.Entry) {
		fmt.Println("evicted", x)
	})
	c.Add("a", 5)
	c.Add("b", 3)
	c.Add("c", 2)
	c.Get("b")
	c.Add("e", 3)
	c.Add("f", 1)

	fmt.Println(c)
}

const (
	ItemsToAdd = 1e7
)

func BenchmarkAny_EncodingBaseline(b *testing.B) {
	b.ResetTimer()
	start := hrtime.Now()
	for k := 0; k < b.N; k++ {
		for i := 0; i < ItemsToAdd; i++ {
			var key [8]byte
			binary.LittleEndian.PutUint64(key[:], uint64(i))
			x := cache.Hash(key[:])
			runtime.KeepAlive(x)
		}
	}
	finish := hrtime.Now()
	b.StopTimer()
	b.ReportMetric(float64((finish-start).Nanoseconds())/float64(b.N*ItemsToAdd), "ns/item")
}

func BenchmarkHalf_Growing(b *testing.B) {
	b.ResetTimer()
	start := hrtime.Now()
	for k := 0; k < b.N; k++ {
		c := cache.NewHalf(1e9, nil)
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

func BenchmarkHalf_Prealloc(b *testing.B) {
	b.ResetTimer()
	start := hrtime.Now()
	for k := 0; k < b.N; k++ {
		c := cache.NewHalfPrealloc(ItemsToAdd, 1e9, nil)
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
