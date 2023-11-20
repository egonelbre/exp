package vector

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

func Apply1To[T any](
	dst, xs Vector[T],
	n uintptr, fn func(x T) T,
) {
	xi, di := xs.Offset, dst.Offset
	for i := uintptr(0); i < n; i++ {
		dst.Values[xi] = fn(xs.Values[xi])
		xi += xs.Inc
		di += dst.Inc
	}
}

func Reduce1[T any](
	xs Vector[T],
	n uintptr,
	zero T,
	fn func(r, x T) T,
) (result T) {
	result = zero
	xi := xs.Offset
	for i := uintptr(0); i < n; i++ {
		result = fn(result, xs.Values[xi])
		xi += xs.Inc
	}
	return result
}

func Apply2[T any](
	xs, ys Vector[T],
	n uintptr,
	fn func(x, y T) T,
) {
	xi, yi := xs.Offset, ys.Offset
	for i := uintptr(0); i < n; i++ {
		xs.Values[i] = fn(xs.Values[xi], ys.Values[yi])
		xi += xs.Inc
		yi += ys.Inc
	}
}

func Apply2To[T any](
	dst, xs, ys Vector[T],
	n uintptr, fn func(x, y T) T,
) {
	di, xi, yi := dst.Offset, xs.Offset, ys.Offset
	for i := uintptr(0); i < n; i++ {
		dst.Values[xi] = fn(xs.Values[xi], ys.Values[yi])
		xi += xs.Inc
		yi += ys.Inc
		di += dst.Inc
	}
}

func Reduce2[T any](
	xs, ys Vector[T],
	n uintptr,
	zero T,
	fn func(r, x, y T) T,
) (result T) {
	result = zero
	xi, yi := xs.Offset, ys.Offset
	for i := uintptr(0); i < n; i++ {
		result = fn(result, xs.Values[xi], ys.Values[xi])
		xi += xs.Inc
		yi += ys.Inc
	}
	return result
}
