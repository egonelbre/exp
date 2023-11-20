package define

import . "github.com/egonelbre/exp/vector/vector"

func AddConst[T float32 | float64 | complex64 | complex128](alpha T, xs []T, n uintptr) {
	Apply1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			return x + alpha
		})
}

func AddConstTo[T float32 | float64 | complex64 | complex128](dst []T, alpha T, xs []T, n uintptr) {
	Apply1To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			return x + alpha
		})
}

func Add[T float32 | float64 | complex64 | complex128](xs, ys []T, n uintptr) {
	Apply2(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x + y
		})
}

func AddTo[T float32 | float64 | complex64 | complex128](dst, xs, ys []T, n uintptr) {
	Apply2To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x + y
		})
}

func Sub[T float32 | float64 | complex64 | complex128](xs, ys []T, n uintptr) {
	Apply2(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x - y
		})
}

func SubTo[T float32 | float64 | complex64 | complex128](dst, xs, ys []T, n uintptr) {
	Apply2To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x - y
		})
}

func Mul[T float32 | float64 | complex64 | complex128](xs, ys []T, n uintptr) {
	Apply2(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x * y
		})
}

func MulTo[T float32 | float64 | complex64 | complex128](dst, xs, ys []T, n uintptr) {
	Apply2To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x * y
		})
}

func Div[T float32 | float64 | complex64 | complex128](xs, ys []T, n uintptr) {
	Apply2(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x * y
		})
}

func DivTo[T float32 | float64 | complex64 | complex128](dst, xs, ys []T, n uintptr) {
	Apply2To(
		Vector[T]{Values: dst, Offset: 0, Inc: 1},
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		Vector[T]{Values: ys, Offset: 0, Inc: 1},
		n, func(x, y T) T {
			return x / y
		})
}
