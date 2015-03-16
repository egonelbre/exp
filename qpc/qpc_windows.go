package qpc

import (
	"syscall"
	"time"
	"unsafe"
)

type Count int64

// precision timing
var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	procFreq    = modkernel32.NewProc("QueryPerformanceFrequency")
	procCounter = modkernel32.NewProc("QueryPerformanceCounter")

	Freq Count = getFrequency()
)

// getFrequency returns frequency in ticks per second
func getFrequency() Count {
	var freq int64
	r1, _, _ := syscall.Syscall(procFreq.Addr(), 1, uintptr(unsafe.Pointer(&freq)), 0, 0)
	if r1 == 0 {
		panic("call failed")
	}
	return Count(freq)
}

// Now returns ticks
func Now() Count {
	var now int64
	r1, _, _ := syscall.Syscall(procCounter.Addr(), 1, uintptr(unsafe.Pointer(&now)), 0, 0)
	if r1 == 0 {
		panic("call failed")
	}
	return Count(now)
}

func (a Count) Sub(b Count) Count { return a - b }

func (count Count) Nanoseconds() int64 {
	return int64(count * 1e9 / Freq)
}

func (count Count) Duration() time.Duration {
	return time.Duration(count) * time.Second / time.Duration(Freq)
}

func Since(start Count) time.Duration { return Now().Sub(start).Duration() }
