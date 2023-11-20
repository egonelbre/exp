package define

import . "github.com/egonelbre/exp/vector/vector"

func Sum[T float32 | float64 | complex64 | complex128](xs []T, n uintptr) T {
	return Reduce1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, 0, func(r, x T) T {
			return r + x
		})
}

func Dot[T float32 | float64 | complex64 | complex128](xs, ys []T, n uintptr) T {
	return Reduce2(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, 0, func(r, x, y T) T {
			return r + x*y
		})
}

func CumSum[T float32 | float64 | complex64 | complex128](xs []T, n uintptr) {
	var result T
	Apply1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			result += x
			return result
		})
}

func CumProd[T float32 | float64 | complex64 | complex128](xs []T, n uintptr) {
	var result T = 1 // TODO: not correct for complex
	Apply1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			result += x
			return result
		})
}
