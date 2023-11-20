//go:build ignore

package define

type Vector[T any] struct {
	Values []T
	Offset uintptr
	Inc    uintptr
}

func Apply1[T any](
	xs Vector[T],
	n uintptr,
	fn func(x T) T,
) {
	xi := xs.Offset
	for i := uintptr(0); i < n; i++ {
		xs.Values[xi] = fn(xs.Values[xi])
		xi += xs.Inc
	}
}

func AddConst[T float32 | float64 | complex64 | complex128](alpha T, xs []T, n uintptr) {
	Apply1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			return x + alpha
		})
}

func AddConst2[T float32 | float64 | complex64 | complex128](alpha T, xs []T, n uintptr) {
	for i := uintptr(0); i < n; i++ {
		xs[i] += alpha
	}
}
