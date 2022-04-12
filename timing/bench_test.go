package timing_test

import (
	"testing"
	"time"

	"github.com/egonelbre/exp/timing"
	"github.com/zeromicro/go-zero/core/collection"
)

func BenchmarkTimingWheel(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		tw, _ := collection.NewTimingWheel(time.Second, 100, func(k, v interface{}) {})
		for k := 0; k < 100000; k++ {
			tw.SetTimer(i, i, time.Second)
			tw.SetTimer(b.N+i, b.N+i, time.Second)
			tw.MoveTimer(i, time.Second*time.Duration(i))
			tw.RemoveTimer(i)
		}
	}
}

func BenchmarkQueue(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q := timing.NewQueue(func(key int, value int) {})
		for k := 0; k < 100000; k++ {
			q.Add(i, i, time.Second)
			q.Add(b.N+i, b.N+i, time.Second)
			q.Move(i, time.Second*time.Duration(i))
			q.Remove(i)
		}
	}
}
