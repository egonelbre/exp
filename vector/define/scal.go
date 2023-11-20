package define

import . "github.com/egonelbre/exp/vector/vector"

func ScalUnitary[T float32 | float64 | complex64 | complex128](alpha T, xs []T, n uintptr) {
	Apply1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			return x * alpha
		})
}

func ScalUnitaryTo[T float32 | float64 | complex64 | complex128](dst []T, alpha T, xs []T, n uintptr) {
	Apply1To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			return x * alpha
		})
}

func ScalIncUnitary[T float32 | float64 | complex64 | complex128](alpha T, xs []T, n, incx uintptr) {
	Apply1(
		Vector[T]{Values: xs, Offset: 0, Inc: incx},
		n, func(x T) T {
			return x * alpha
		})
}

func ScalIncUnitaryTo[T float32 | float64 | complex64 | complex128](dst []T, incdst uintptr, alpha T, xs []T, n, incx uintptr) {
	Apply1To(
		Vector[T]{Values: dst, Offset: 0, Inc: incdst},
		Vector[T]{Values: xs, Offset: 0, Inc: incx},
		n, func(x T) T {
			return x * alpha
		})
}
