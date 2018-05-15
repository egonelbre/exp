package audiobuffer

import "testing"

type Interleaved struct{ Data []float32 }

func NewInterleaved(frames int) *Interleaved {
	return &Interleaved{
		Data: make([]float32, frames*2),
	}
}

type Interleaved2 struct{ Data [][2]float32 }

func NewInterleaved2(frames int) *Interleaved2 {
	return &Interleaved2{
		Data: make([][2]float32, frames),
	}
}

type Deinterleaved struct{ Data []float32 }

func NewDeinterleaved(frames int) *Deinterleaved {
	return &Deinterleaved{
		Data: make([]float32, frames*2),
	}
}

const BufferSize = 64

func params() (a, b float32) {
	return 0.5, 0.25
}

func BenchmarkInterleaved(b *testing.B) {
	buf := NewInterleaved(BufferSize)
	for i := 0; i < b.N; i++ {
		a, b := params()
		for i := 0; i < len(buf.Data); i += 2 {
			buf.Data[i+0] *= a
			buf.Data[i+1] *= b
			a *= 0.99999
			b *= 0.99999
		}
	}
}

func BenchmarkInterleavedVariable(b *testing.B) {
	buf := NewInterleaved(BufferSize)
	n := 2
	for i := 0; i < b.N; i++ {
		a, b := params()
		s := []float32{a, b}
		for i := 0; i < len(buf.Data); i += n {
			for k := 0; k < n; k++ {
				buf.Data[i+k] *= s[k]
				s[k] *= 0.99999
			}
		}
	}
}

func BenchmarkInterleaved2(b *testing.B) {
	buf := NewInterleaved2(BufferSize)
	for i := 0; i < b.N; i++ {
		a, b := params()
		for i := range buf.Data {
			buf.Data[i][0] *= a
			buf.Data[i][1] *= b
			a *= 0.99999
			b *= 0.99999
		}
	}
}

func BenchmarkDeinterleaved(b *testing.B) {
	buf := NewDeinterleaved(BufferSize)
	for i := 0; i < b.N; i++ {
		a, b := params()
		left := buf.Data[:len(buf.Data)/2]
		for i := range left {
			left[i] *= a
			a *= 0.99999
		}
		right := buf.Data[len(buf.Data)/2:]
		for i := range right {
			right[i] *= b
			b *= 0.99999
		}
	}
}
