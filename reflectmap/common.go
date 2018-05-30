package reflectmap

import "unsafe"

const bucketSize = 8

type hashBucket [bucketSize]byte

type reflectvalue struct {
	typ uintptr
	ptr unsafe.Pointer
}
