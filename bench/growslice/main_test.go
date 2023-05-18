// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math/bits"
	"testing"
	"unsafe"
)

var bytesType = &_type{Size_: 1, PtrBytes: 0}
var intsType = &_type{Size_: goarch_PtrSize, PtrBytes: 0}
var uint32Type = &_type{Size_: 4, PtrBytes: 0}
var ptrType = &_type{Size_: goarch_PtrSize, PtrBytes: goarch_PtrSize}
var sliceType = &_type{Size_: 3 * goarch_PtrSize, PtrBytes: goarch_PtrSize}
var ifaceType = &_type{Size_: 2 * goarch_PtrSize, PtrBytes: 2 * goarch_PtrSize}

func BenchmarkGrowslice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		growslice(nil, 1232, 64, 1232-48, bytesType)
		growslice(nil, 68, 64, 68-48, ptrType)
		growslice(nil, 123, 32, 123-48, intsType)
		growslice(nil, 8, 4, 8-6, sliceType)
		growslice(nil, 12, 8, 12-48, uint32Type)
		growslice(nil, 1232, 64, 1232-48, ifaceType)
	}
}

//go:noinline
func growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type) slice {
	oldLen := newLen - num
	/*
		if raceenabled {
			callerpc := getcallerpc()
			racereadrangepc(oldPtr, uintptr(oldLen*int(et.Size_)), callerpc, abi.FuncPCABIInternal(growslice))
		}
		if msanenabled {
			msanread(oldPtr, uintptr(oldLen*int(et.Size_)))
		}
		if asanenabled {
			asanread(oldPtr, uintptr(oldLen*int(et.Size_)))
		}
	*/

	if newLen < 0 {
		panic(errorString("growslice: len out of range"))
	}

	if et.Size_ == 0 {
		// append should not create a slice with nil pointer but non-zero len.
		// We assume that append doesn't need to preserve oldPtr in this case.
		return slice{unsafe.Pointer(&zerobase), newLen, newLen}
	}

	var newcap int
	if enable_nextslicecap {
		newcap = nextslicecap(newLen, oldCap)
	} else {
		newcap = oldCap
		doublecap := newcap + newcap
		if newLen > doublecap {
			newcap = newLen
		} else {
			const threshold = 256
			if oldCap < threshold {
				newcap = doublecap
			} else {
				// Check 0 < newcap to detect overflow
				// and prevent an infinite loop.
				for 0 < newcap && newcap < newLen {
					// Transition from growing 2x for small slices
					// to growing 1.25x for large slices. This formula
					// gives a smooth-ish transition between the two.
					newcap += (newcap + 3*threshold) / 4
				}
				// Set newcap to the requested cap when
				// the newcap calculation overflowed.
				if newcap <= 0 {
					newcap = newLen
				}
			}
		}
	}

	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.Size.
	// For 1 we don't need any division/multiplication.
	// For goarch_PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	switch {
	case et.Size_ == 1:
		lenmem = uintptr(oldLen)
		newlenmem = uintptr(newLen)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.Size_ == goarch_PtrSize:
		lenmem = uintptr(oldLen) * goarch_PtrSize
		newlenmem = uintptr(newLen) * goarch_PtrSize
		capmem = roundupsize(uintptr(newcap) * goarch_PtrSize)
		overflow = uintptr(newcap) > maxAlloc/goarch_PtrSize
		newcap = int(capmem / goarch_PtrSize)
	case isPowerOfTwo(et.Size_):
		var shift uintptr
		if goarch_PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(bits.TrailingZeros64(uint64(et.Size_))) & 63
		} else {
			shift = uintptr(bits.TrailingZeros32(uint32(et.Size_))) & 31
		}
		lenmem = uintptr(oldLen) << shift
		newlenmem = uintptr(newLen) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
		capmem = uintptr(newcap) << shift
	default:
		lenmem = uintptr(oldLen) * et.Size_
		newlenmem = uintptr(newLen) * et.Size_
		capmem, overflow = math_MulUintptr(et.Size_, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.Size_)
		capmem = uintptr(newcap) * et.Size_
	}

	// The check of overflow in addition to capmem > maxAlloc is needed
	// to prevent an overflow which can be used to trigger a segfault
	// on 32bit architectures with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	//   s = append(s, d, d, d, d)
	//   print(len(s), "\n")
	// }
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: len out of range"))
	}

	var p unsafe.Pointer
	if et.PtrBytes == 0 {
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from oldLen to newLen.
		// Only clear the part that will not be overwritten.
		// The reflect_growslice() that calls growslice will manually clear
		// the region not cleared here.
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			// Only shade the pointers in oldPtr since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(oldPtr), lenmem-et.Size_+et.PtrBytes)
		}
	}
	memmove(p, oldPtr, lenmem)

	return slice{p, newLen, newcap}
}

// nextslicecap computes the next appropriate slice length.
func nextslicecap(newLen, oldCap int) int {
	newcap := oldCap
	doublecap := newcap + newcap
	if newLen > doublecap {
		return newLen
	}

	const threshold = 256
	if oldCap < threshold {
		return doublecap
	}
	for {
		// Transition from growing 2x for small slices
		// to growing 1.25x for large slices. This formula
		// gives a smooth-ish transition between the two.
		newcap += (newcap + 3*threshold) >> 2

		// We need to check `newcap >= newLen` and whether `newcap` overflowed.
		// newLen is guaranteed to be larger than zero, hence
		// when newcap overflows then `uint(newcap) > uint(newLen)`.
		// This allows to check for both with the same comparison.
		if uint(newcap) >= uint(newLen) {
			break
		}
	}

	// Set newcap to the requested cap when
	// the newcap calculation overflowed.
	if newcap <= 0 {
		return newLen
	}
	return newcap
}
