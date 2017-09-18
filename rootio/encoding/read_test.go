package hep

import (
	"testing"
)

// Naming
//   B / L = big endian / little endian
//   A / S = (at, data) / (data[at:])
//   F / R = forward vs reverse
//   U = no bounds checks

func BenchmarkReadU32(b *testing.B) {
	data := make([]byte, 1003)
	for i := range data {
		data[i] = byte(i)
	}
	b.Run("binary.BigEndian", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += BigU32(data[at:])
			}
		}
	})
	b.Run("BAR", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += BigU32ar(at, data)
			}
		}
	})
	b.Run("BAF", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += BigU32af(at, data)
			}
		}
	})
	b.Run("LAF", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += LittleU32af(at, data)
			}
		}
	})
	b.Run("LAR", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += LittleU32ar(at, data)
			}
		}
	})
	b.Run("NAD", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += NativeU32ad(at, data)
			}
		}
	})
	b.Run("NADU", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += NativeU32adu(at, data)
			}
		}
	})
	b.Run("BSR", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += BigU32sr(data[at:])
			}
		}
	})
	b.Run("BSF", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += BigU32sf(data[at:])
			}
		}
	})
	b.Run("binary.LittleEndian", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += LittleU32(data[at:])
			}
		}
	})
	b.Run("LSF", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += LittleU32sf(data[at:])
			}
		}
	})
	b.Run("LSR", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += LittleU32sr(data[at:])
			}
		}
	})
	b.Run("NSD", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += NativeU32sd(data[at:])
			}
		}
	})
	b.Run("NSDU", func(b *testing.B) {
		t := uint32(0)
		for i := 0; i < b.N; i++ {
			for at := 0; at < len(data)-4; at++ {
				t += NativeU32sdu(data[at:])
			}
		}
	})
}
