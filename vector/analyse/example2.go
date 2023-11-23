//go:build ignore

package define

func Axpy(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	var xi, yi uintptr
	for i := uintptr(0); i < n; i++ {
		ys[yi] += alpha * xs[xi]
		xi += incx
		yi += incy
	}
}
