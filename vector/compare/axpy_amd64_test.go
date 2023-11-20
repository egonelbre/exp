package vecprocess

import "testing"

func BenchmarkAxpyAssembly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AxpyInc(alpha, xs[i&3:], ys, K, 2, 4, 0, 0)
	}
}
