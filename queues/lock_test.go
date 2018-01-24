// implements fifo using a channel
package goqueuestest
/*
import (
	"sync"
	"testing"
)

func benchLockChan(b *testing.B)

func BenchmarkLockChan(b *testing.B) {
	ch := make(chan int, 1)
	ch <- 1
	runParallelBench(b, 8, func(q Queue, N int){
		for i := 0; i < N; i += 1 {
			<- ch
			ch <- 1
		}
	})
}

func BenchmarkLockMutex(b *testing.B) {
	m := sync.Mutex{}
	runParallelBench(b, func(q Queue, N int){
		for i := 0; i < N; i += 1 {
			m.Lock()
			m.Unlock()
		}
	})
}*/