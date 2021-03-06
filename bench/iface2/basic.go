package main

import (
	"fmt"
	"runtime"

	"github.com/loov/hrtime"
)

func main() {
	runtime.LockOSThread()

	const N = 100000000

	var overhead uint64 = 100
	for i := 0; i < N; i++ {
		start := hrtime.RDTSC()
		stop := hrtime.RDTSCP()
		delta := stop - start
		if delta < overhead {
			overhead = delta
		}
	}

	averageOverhead := float64(overhead)
	fmt.Println("measurement overhead:", averageOverhead, "cy")

	{
		var measurement uint64 = 100
		for i := 0; i < N; i++ {
			start := hrtime.RDTSC()

			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())

			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())
			runtime.KeepAlive(example.Get())

			stop := hrtime.RDTSCP()
			delta := stop - start
			if delta < measurement {
				measurement = delta
			}

			if i & 1 == 1 {
				example = &A{}
			} else {
				example = &B{}
			}
		}

		averageCall := float64(measurement)
		fmt.Println("iface call:", float64(averageCall-averageOverhead) / 10, "cy")
	}


	{
		var measurement uint64 = 100
		for i := 0; i < N; i++ {
			start := hrtime.RDTSC()
			
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())

			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())
			runtime.KeepAlive(nop())

			stop := hrtime.RDTSCP()
			delta := stop - start
			if delta < measurement {
				measurement = delta
			}

			if i & 1 == 1 {
				example = &A{}
			} else {
				example = &B{}
			}
		}

		averageCall := float64(measurement)
		fmt.Println("nop call:", float64(averageCall-averageOverhead) / 10, "cy")
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