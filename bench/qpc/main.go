package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

// precision timing
var (
	modkernel32                   = syscall.NewLazyDLL("kernel32.dll")
	procQueryPerformanceFrequency = modkernel32.NewProc("QueryPerformanceFrequency")
	procQueryPerformanceCounter   = modkernel32.NewProc("QueryPerformanceCounter")
)

func asmQPC() int64

var (
	_QueryPerformanceFrequency uintptr
	_QueryPerformanceCounter   uintptr
)

type stdFunction unsafe.Pointer

// now returns time.Duration using queryPerformanceCounter
func QPC() int64 {
	var now int64
	syscall.Syscall(procQueryPerformanceCounter.Addr(), 1, uintptr(unsafe.Pointer(&now)), 0, 0)
	return now
}

// now returns time.Duration using queryPerformanceCounter
func QPC2() int64 {
	var now int64
	syscall.SyscallN(procQueryPerformanceCounter.Addr(), uintptr(unsafe.Pointer(&now)))
	return now
}

// QPCFrequency returns frequency in ticks per second
func QPCFrequency() int64 {
	var freq int64
	r1, _, _ := syscall.Syscall(procQueryPerformanceFrequency.Addr(), 1, uintptr(unsafe.Pointer(&freq)), 0, 0)
	if r1 == 0 {
		panic("call failed")
	}
	return freq
}

func main() {
	_QueryPerformanceFrequency = procQueryPerformanceFrequency.Addr()
	_QueryPerformanceCounter = procQueryPerformanceCounter.Addr()
	const N = 10000000

	{
		start := QPC()
		for i := 0; i < N; i++ {
			QPC()
		}
		finish := QPC()
		fmt.Println("QPC", delta(finish, start)/N)
	}

	{
		start := QPC()
		for i := 0; i < N; i++ {
			QPC2()
		}
		finish := QPC()
		fmt.Println("QPC2", delta(finish, start)/N)
	}

	{
		start := QPC()
		for i := 0; i < N; i++ {
			asmQPC()
		}

		finish := QPC()
		fmt.Println("asmQPC", delta(finish, start)/N)
		fmt.Println("check", asmQPC()-QPC())
	}

	{
		start := QPC()
		for i := 0; i < N; i++ {
			time.Now()
		}
		finish := QPC()
		fmt.Println("time.Now", delta(finish, start)/N)
	}
}

func delta(a, b int64) time.Duration {
	return time.Duration(a-b) * time.Second / (time.Duration(QPCFrequency()) * time.Nanosecond)
}
