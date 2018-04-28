package queue

import (
	"sync/atomic"
	"unsafe"
)

func ceil(a, n int) int {
	r := ((a + n - 1) / n) * n
	if r <= n {
		return n * 2
	}
	return r
}

func exchangePointer(addr *unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	for {
		prev := atomic.LoadPointer(addr)
		if atomic.CompareAndSwapPointer(addr, prev, value) {
			return prev
		}
	}
}
