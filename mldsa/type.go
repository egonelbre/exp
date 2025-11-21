package mldsa

import (
	simd "simd/archsimd"
	"unsafe"
)

// Uint32x8 type helpers

func BroadcastUint32x8[T ~uint32](v T) simd.Uint32x8 {
	return simd.BroadcastUint32x8(uint32(v))
}

func LoadUint32x8Slice[T ~uint32](at []T) simd.Uint32x8 {
	return LoadUint32x8Pointer(unsafe.SliceData(at))
}

func StoreUint32x8Slice[T ~uint32](v simd.Uint32x8, at []T) {
	StoreUint32x8Pointer(v, unsafe.SliceData(at))
}

func LoadUint32x8Pointer[T ~uint32](at *T) simd.Uint32x8 {
	ptr := (*[8]uint32)(unsafe.Pointer(at))
	return simd.LoadUint32x8(ptr)
}

func StoreUint32x8Pointer[T ~uint32](v simd.Uint32x8, at *T) {
	ptr := (*[8]uint32)(unsafe.Pointer(at))
	v.Store(ptr)
}

func LoadUint32x8[T ~uint32](at *[8]T) simd.Uint32x8 {
	ptr := (*[8]uint32)(unsafe.Pointer(at))
	return simd.LoadUint32x8(ptr)
}

func StoreUint32x8[T ~uint32](v simd.Uint32x8, at *[8]T) {
	ptr := (*[8]uint32)(unsafe.Pointer(at))
	v.Store(ptr)
}
