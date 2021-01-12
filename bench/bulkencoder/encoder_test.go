// Hopefully there are no bugs.

// RasperryPi 4
//
// goos: linux
// goarch: arm64
// BenchmarkBulk-4            19346             62017 ns/op
// BenchmarkByte-4            13394             89582 ns/op
package encoder_test

import (
	"encoding/binary"
	"testing"
)

type BulkEncoder struct {
	fill uint64
	off  uint32

	at   uint64
	data []uint64
}

func NewBulkEncoder(bytes int) *BulkEncoder {
	return &BulkEncoder{data: make([]uint64, bytes/8)}
}

func (enc *BulkEncoder) WriteByte(v byte) {
	enc.fill |= uint64(v) << enc.off
	enc.off += 8
	if enc.off >= 64 {
		enc.data[enc.at] = enc.fill
		enc.at++
		enc.fill = 0
		enc.off = 0
	}
}

func (enc *BulkEncoder) WriteUint32(v uint32) {
	enc.fill |= uint64(v) << enc.off
	enc.off += 32
	if enc.off >= 64 {
		enc.off -= 64
		enc.fill = uint64(v) >> (32 - enc.off)
		enc.data[enc.at] = enc.fill
		enc.at++
	}
}

func (enc *BulkEncoder) Reset() {
	enc.fill, enc.off, enc.at = 0, 0, 0
}

type ByteEncoder struct {
	at   uint64
	data []byte
}

func NewByteEncoder(bytes int) *ByteEncoder {
	return &ByteEncoder{data: make([]byte, bytes)}
}

func (enc *ByteEncoder) WriteByte(v byte) {
	enc.data[enc.at] = v
	enc.at++
}

func (enc *ByteEncoder) WriteUint32(v uint32) {
	binary.LittleEndian.PutUint32(enc.data[enc.at:], v)
	enc.at += 4
}

func (enc *ByteEncoder) Reset() { enc.at = 0 }

const N = 10 * 1024

func BenchmarkBulk_byte(b *testing.B) {
	enc := NewBulkEncoder(N)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteByte(byte(i))
		}
	}
}

func BenchmarkBulk_4xbyte(b *testing.B) {
	enc := NewBulkEncoder(4 * N)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteByte(byte(i))
			enc.WriteByte(byte(i))
			enc.WriteByte(byte(i))
			enc.WriteByte(byte(i))
		}
	}
}

func BenchmarkBulk_uint32(b *testing.B) {
	enc := NewBulkEncoder(N * 4)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteUint32(uint32(i))
		}
	}
}

func BenchmarkBulk_4xuint32(b *testing.B) {
	enc := NewBulkEncoder(N * 4 * 4)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteUint32(uint32(i))
			enc.WriteUint32(uint32(i))
			enc.WriteUint32(uint32(i))
			enc.WriteUint32(uint32(i))
		}
	}
}

func BenchmarkBulk_mix(b *testing.B) {
	enc := NewBulkEncoder(N + N*4 + 32)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteByte(byte(i))
			enc.WriteUint32(uint32(i))
		}
	}
}

func BenchmarkByte_byte(b *testing.B) {
	enc := NewByteEncoder(N)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteByte(byte(i))
		}
	}
}

func BenchmarkByte_4xbyte(b *testing.B) {
	enc := NewByteEncoder(4 * N)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteByte(byte(i))
			enc.WriteByte(byte(i))
			enc.WriteByte(byte(i))
			enc.WriteByte(byte(i))
		}
	}
}

func BenchmarkByte_uint32(b *testing.B) {
	enc := NewByteEncoder(N * 4)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteUint32(uint32(i))
		}
	}
}

func BenchmarkByte_4xuint32(b *testing.B) {
	enc := NewByteEncoder(N * 4 * 4)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteUint32(uint32(i))
			enc.WriteUint32(uint32(i))
			enc.WriteUint32(uint32(i))
			enc.WriteUint32(uint32(i))
		}
	}
}

func BenchmarkByte_mix(b *testing.B) {
	enc := NewByteEncoder(N + N*4 + 32)
	for k := 0; k < b.N; k++ {
		enc.Reset()
		for i := 0; i < N; i++ {
			enc.WriteByte(byte(i))
			enc.WriteUint32(uint32(i))
		}
	}
}
