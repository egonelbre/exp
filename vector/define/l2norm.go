package define

import . "github.com/egonelbre/exp/vector/vector"

func L2NormUnitary[T float32 | float64 | complex64 | complex128](xs []T, n uintptr) T {
	var sumSquares T = 1
	var scale T

	return Reduce1(
		Vector[T]{Values: xs, Offset: 0, Inc: 1},
		n, func(x T) T {
			if x == 0 {

			}
			result += x
			return result
		})
}
