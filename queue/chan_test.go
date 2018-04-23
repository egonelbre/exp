// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package queue

import (
	"runtime"
	"sync/atomic"
	"testing"
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

func BenchmarkChanProdCons10(b *testing.B) {
	benchmarkChanProdCons(b, 10, 0)
}

func BenchmarkChanProdCons100(b *testing.B) {
	benchmarkChanProdCons(b, 100, 0)
}

func BenchmarkChanProdConsWork10(b *testing.B) {
	benchmarkChanProdCons(b, 10, 100)
}

func BenchmarkChanProdConsWork100(b *testing.B) {
	benchmarkChanProdCons(b, 100, 100)
}
