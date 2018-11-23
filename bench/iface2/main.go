package main

import (
	"fmt"
	"runtime"

	"github.com/loov/hrtime"
)

func main() {
	runtime.LockOSThread()

	const N = 100000000
	const ProbablyGC = 100

	var overhead uint64
	var overheadCount int
	for i := 0; i < N; i++ {
		start := hrtime.RDTSC()
		stop := hrtime.RDTSCP()
		delta := stop - start
		if delta < ProbablyGC {
			overhead += delta
			overheadCount++
		}
	}

	averageOverhead := float64(overhead) / float64(overheadCount)
	fmt.Println("measurement overhead:", averageOverhead, "cy")

	{
		var measurement uint64
		var measurementCount int
		for i := 0; i < N; i++ {
			start := hrtime.RDTSC()
			x := example.Get()
			runtime.KeepAlive(x)
			stop := hrtime.RDTSCP()
			delta := stop - start
			if delta < ProbablyGC {
				measurement += delta
				measurementCount++
			}

			if i & 1 == 1 {
				example = &A{}
			} else {
				example = &B{}
			}
		}

		averageCall := float64(measurement) / float64(measurementCount)
		fmt.Println("iface call:", averageCall-averageOverhead, "cy")
	}


	{
		var measurement uint64
		var measurementCount int
		for i := 0; i < N; i++ {
			start := hrtime.RDTSC()
			nop()
			stop := hrtime.RDTSCP()
			delta := stop - start
			if delta < ProbablyGC {
				measurement += delta
				measurementCount++
			}

			if i & 1 == 1 {
				example = &A{}
			} else {
				example = &B{}
			}
		}

		averageCall := float64(measurement) / float64(measurementCount)
		fmt.Println("nop call:", averageCall-averageOverhead, "cy")
	}
}

var example Interface = &A{}

type Interface interface {
	Get() int
}

type A struct{ _ int; Value int }
type B struct{ _ int; Value int }

func (x *A) Get() int { return x.Value }
func (x *B) Get() int { return x.Value }


//go:noinline
func nop() int { return 0 }