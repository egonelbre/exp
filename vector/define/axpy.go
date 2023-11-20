package define

import . "github.com/egonelbre/exp/vector/vector"

// AxpyUnitary calculates y[i] = y[i] + alpha * x[i].
func AxpyUnitary[T float32 | float64 | complex64 | complex128](alpha T, xs, ys []T) {
	Apply2(
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		uintptr(len(ys)), func(y, x T) T {
			return y + alpha*x
		})
}

// AxpyUnitaryTo calculates dst[i] = y[i] + alpha * x[i].
func AxpyUnitaryTo[T float32 | float64 | complex64 | complex128](dst []T, alpha T, xs, ys []T) {
	Apply2To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		uintptr(len(ys)), func(y, x T) T {
			return y + x*alpha
		})
}
