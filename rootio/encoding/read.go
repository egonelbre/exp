package hep

import (
	"encoding/binary"
	"unsafe"
)

func LittleU32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

func BigU32ar(at int, data []byte) uint32 {
	return uint32(data[at+3]) | uint32(data[at+2])<<8 | uint32(data[at+1])<<16 | uint32(data[at])<<24
}

func BigU32af(at int, data []byte) uint32 {
	return uint32(data[at])<<24 | uint32(data[at+1])<<16 | uint32(data[at+2])<<8 | uint32(data[at+3])
}

func LittleU32af(at int, data []byte) uint32 {
	return uint32(data[at]) | uint32(data[at+1])<<8 | uint32(data[at+2])<<16 | uint32(data[at+3])<<24
}

func LittleU32ar(at int, data []byte) uint32 {
	return uint32(data[at+3])<<24 | uint32(data[at+2])<<16 | uint32(data[at+1])<<8 | uint32(data[at])
}

func NativeU32ad(at int, data []byte) uint32 {
	_ = data[3]
	return *(*uint32)(unsafe.Pointer(&data[at]))
}
func NativeU32adu(at int, data []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&data[at]))
}

func BigU32(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

func BigU32sr(data []byte) uint32 {
	return uint32(data[3]) | uint32(data[2])<<8 | uint32(data[1])<<16 | uint32(data[0])<<24
}

func BigU32sf(data []byte) uint32 {
	return uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
}

func LittleU32sf(data []byte) uint32 {
	return uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
}

func LittleU32sr(data []byte) uint32 {
	return uint32(data[3])<<24 | uint32(data[2])<<16 | uint32(data[0+1])<<8 | uint32(data[0])

}

func NativeU32sd(data []byte) uint32 {
	_ = data[3]
	return *(*uint32)(unsafe.Pointer(&data[0]))
}

func NativeU32sdu(data []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&data[0]))
}
