package compare

import "testing"

const N = 10000000
const K = N/5 - 1

var alpha = float32(1.0)
var xs = makeFloat32(N)
var ys = makeFloat32(N)

func BenchmarkAxpyBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyBasic(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyUnsafe(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyUnsafeInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyUnsafeInline(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyPointer(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyPointerLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyPointerLoop(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyBasicR4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyBasicR4(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyUnsafeR4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyUnsafeR4(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyUnsafeInlineR4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyUnsafeInlineR4(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyUnsafeInlineR8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyUnsafeInlineR8(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyPointerR4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyPointerR4(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyPointerLoopR4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyPointerLoopR4(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyPointerLoopInterleaveR4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyPointerLoopInterleaveR4(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func BenchmarkAxpyPointerR4Alt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyPointerR4Alt(alpha, xs[i&3:], 2, ys, 4, K)
	}
}

func makeFloat32(n int) []float32 {
	r := make([]float32, n)
	for i := range r {
		r[i] = float32(i)
	}
	return r
}
