// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package queue

import (
	"runtime"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/egonelbre/async"
)

func BenchmarkChanUncontended(b *testing.B) {
	const C = 100
	b.RunParallel(func(pb *testing.PB) {
		myc := make(chan int, C)
		for pb.Next() {
			for i := 0; i < C; i++ {
				myc <- 0
			}
			for i := 0; i < C; i++ {
				<-myc
			}
		}
	})
}

func BenchmarkChanTryRecv(b *testing.B) {
	const C = 100
	myc := make(chan int, C*runtime.GOMAXPROCS(0))
	b.RunParallel(func(pb *testing.PB) {
		var value int
		for pb.Next() {
			select {
			case myc <- value:
			default:
			}
			select {
			case value = <-myc:
			default:
			}
		}
	})
}

func BenchmarkChanContended(b *testing.B) {
	const C = 100
	myc := make(chan int, C*runtime.GOMAXPROCS(0))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < C; i++ {
				myc <- 0
			}
			for i := 0; i < C; i++ {
				<-myc
			}
		}
	})
}
func benchmarkChanProdCons(b *testing.B, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := int32(b.N / CallsPerSched)
	c := make(chan bool, 2*procs)
	myc := make(chan int, chanSize)
	for p := 0; p < procs; p++ {
		go func() {
			foo := 0
			for atomic.AddInt32(&N, -1) >= 0 {
				for g := 0; g < CallsPerSched; g++ {
					for i := 0; i < localWork; i++ {
						foo *= 2
						foo /= 2
					}
					myc <- 1
				}
			}
			myc <- 0
			c <- foo == 42
		}()
		go func() {
			foo := 0
			for {
				v := <-myc
				if v == 0 {
					break
				}
				for i := 0; i < localWork; i++ {
					foo *= 2
					foo /= 2
				}
			}
			c <- foo == 42
		}()
	}
	for p := 0; p < procs; p++ {
		<-c
		<-c
	}
}

func BenchmarkChanProdCons(b *testing.B) {
	for _, chanSize := range []int{1, 10, 100} {
		b.Run(strconv.Itoa(chanSize), func(b *testing.B) {
			benchmarkChanProdCons(b, chanSize, 0)
		})
	}
	for _, chanSize := range []int{1, 10, 100} {
		b.Run("Work"+strconv.Itoa(chanSize), func(b *testing.B) {
			benchmarkChanProdCons(b, chanSize, 100)
		})
	}
}

func BenchmarkChanAsymmetric(b *testing.B) {
	const C = 5
	const N = 200
	for k := 0; k < b.N; k++ {
		q := make(chan int, N*C)
		async.Spawn(C, func(int) {
			for i := 0; i < N; i++ {
				q <- i
			}
		})

		for i := 0; i < C*N; i++ {
			_ = <-q
		}
	}
}
