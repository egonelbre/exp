// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "unsafe"

// base address for all 0-byte allocations
var zerobase uintptr

type errorString string

func (e errorString) RuntimeError() {}

func (e errorString) Error() string {
	return "runtime error: " + string(e)
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type Type struct {
	Size_       uintptr
	PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
	Hash        uint32  // hash of type; avoids computation in hash tables
	TFlag       uint8   // extra type information flags
	Align_      uint8   // alignment of variable with this type
	FieldAlign_ uint8   // alignment of struct field with this type
	Kind_       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
	// GCData stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, GCData is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	GCData    *byte
	Str       int32 // string form
	PtrToThis int32 // type for pointer to this type, may be zero
}

type _type = Type

//go:noinline
func memmove(p, oldptr unsafe.Pointer, lenmem uintptr) {}

//go:nosplit
//go:noinline
func bulkBarrierPreWriteSrcOnly(dst, src, size uintptr) {}

//go:noinline
func memclrNoHeapPointersChunked(size uintptr, x unsafe.Pointer) {}

//go:noinline
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer { return nil }

var writeBarrier struct {
	enabled bool    // compiler emits a check of this before calling write barrier
	pad     [3]byte // compiler uses 32-bit load for "enabled" field
	needed  bool    // identical to enabled, for now (TODO: dedup)
	alignme uint64  // guarantee alignment so that compiler can use a 32 or 64-bit load
}

//go:noinline
func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr) {}

//go:noinline
func add(p unsafe.Pointer, newlenmem uintptr) unsafe.Pointer {
	return nil
}

// this is actually a compiler intrinsic
func abi_FuncPCABIInternal(f interface{}) uintptr { return 0 }

//go:noinline
func getcallerpc() uintptr { return 0 }

//go:noinline
func racereadrangepc(addr unsafe.Pointer, sz, callerpc, pc uintptr) {}

//go:noinline
func msanread(addr unsafe.Pointer, sz uintptr) {}

//go:noinline
func asanread(addr unsafe.Pointer, sz uintptr) {}
