package compare

import (
	"unsafe"
)

func at[T any](xs []T, index uintptr) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(xs)), index*unsafe.Sizeof(xs[0])))
}

//go:noinline
func AxpyBasic(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	for i := uintptr(0); i < n; i++ {
		ys[yi] += alpha * xs[xi]

		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafe(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	for i := uintptr(0); i < n; i++ {
		*at(ys, yi) += alpha * *at(xs, xi)
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeX(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	for ; n > 0; n-- {
		*at(ys, yi) += alpha * *at(xs, xi)
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeInline(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	for i := uintptr(0); i < n; i++ {
		*at(ys, i*incy) += alpha * *at(xs, i*incx)
	}
}

//go:noinline
func AxpyPointer(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))
	xn := unsafe.Add(xp, 4*n*incx)
	for uintptr(xp) < uintptr(xn) {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, 4*incx), unsafe.Add(yp, 4*incy)
	}
}

//go:noinline
func AxpyPointerLoop(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))
	for i := uintptr(0); i < n; i++ {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, 4*incx), unsafe.Add(yp, 4*incy)
	}
}

//go:noinline
func AxpyPointerLoopX(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))
	for ; n > 0; n-- {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, 4*incx), unsafe.Add(yp, 4*incy)
	}
}

const mask4 = ^uintptr(3)
const mask8 = ^uintptr(7)

//go:noinline
func AxpyBasicR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	i := uintptr(0)
	n4 := n & mask4
	for ; i < n4; i += 4 {
		ys[yi+0] += alpha * xs[xi+0]
		ys[yi+1] += alpha * xs[xi+1]
		ys[yi+2] += alpha * xs[xi+2]
		ys[yi+3] += alpha * xs[xi+3]

		xi += incx * 4
		yi += incy * 4
	}
	for ; i < n; i++ {
		ys[yi] += alpha * xs[xi]
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	i := uintptr(0)
	n4 := n & mask4
	for ; i < n4; i += 4 {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		*at(ys, yi+1) += alpha * *at(xs, xi+1)
		*at(ys, yi+2) += alpha * *at(xs, xi+2)
		*at(ys, yi+3) += alpha * *at(xs, xi+3)
		xi += incx * 4
		yi += incy * 4
	}
	for ; i < n; i++ {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyBasicXR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	for ; n >= 4; n -= 4 {
		ys[yi+0] += alpha * xs[xi+0]
		ys[yi+1] += alpha * xs[xi+1]
		ys[yi+2] += alpha * xs[xi+2]
		ys[yi+3] += alpha * xs[xi+3]

		xi += incx * 4
		yi += incy * 4
	}
	for ; n > 0; n-- {
		ys[yi] += alpha * xs[xi]
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeXR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	xi, yi := uintptr(0), uintptr(0)
	for ; n >= 4; n -= 4 {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		*at(ys, yi+1) += alpha * *at(xs, xi+1)
		*at(ys, yi+2) += alpha * *at(xs, xi+2)
		*at(ys, yi+3) += alpha * *at(xs, xi+3)
		xi += incx * 4
		yi += incy * 4
	}
	for ; n > 0; n-- {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeR8(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	if n == 0 {
		return
	}
	_, _ = xs[n*incx-1], ys[n*incy-1]
	xi, yi := uintptr(0), uintptr(0)
	n8 := n & mask8
	i := uintptr(0)
	for ; i < n8; i += 8 {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		*at(ys, yi+1) += alpha * *at(xs, xi+1)
		*at(ys, yi+2) += alpha * *at(xs, xi+2)
		*at(ys, yi+3) += alpha * *at(xs, xi+3)
		*at(ys, yi+4) += alpha * *at(xs, xi+4)
		*at(ys, yi+5) += alpha * *at(xs, xi+5)
		*at(ys, yi+6) += alpha * *at(xs, xi+6)
		*at(ys, yi+7) += alpha * *at(xs, xi+7)
		xi += incx * 8
		yi += incy * 8
	}
	for ; i < n; i++ {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeXR8(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	if n == 0 {
		return
	}
	_, _ = xs[n*incx-1], ys[n*incy-1]
	xi, yi := uintptr(0), uintptr(0)
	for ; n >= 8; n -= 8 {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		*at(ys, yi+1) += alpha * *at(xs, xi+1)
		*at(ys, yi+2) += alpha * *at(xs, xi+2)
		*at(ys, yi+3) += alpha * *at(xs, xi+3)
		*at(ys, yi+4) += alpha * *at(xs, xi+4)
		*at(ys, yi+5) += alpha * *at(xs, xi+5)
		*at(ys, yi+6) += alpha * *at(xs, xi+6)
		*at(ys, yi+7) += alpha * *at(xs, xi+7)
		xi += incx * 8
		yi += incy * 8
	}
	for ; n > 0; n-- {
		*at(ys, yi+0) += alpha * *at(xs, xi+0)
		xi += incx
		yi += incy
	}
}

//go:noinline
func AxpyUnsafeInlineR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	i := uintptr(0)
	n4 := n & mask4
	for ; i < n4; i += 4 {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
		*at(ys, (i+1)*incy) += alpha * *at(xs, (i+1)*incx)
		*at(ys, (i+2)*incy) += alpha * *at(xs, (i+2)*incx)
		*at(ys, (i+3)*incy) += alpha * *at(xs, (i+3)*incx)
	}
	for ; i < n; i++ {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
	}
}

//go:noinline
func AxpyUnsafeInlineXR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	i := uintptr(0)
	for ; n >= 4; n -= 4 {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
		*at(ys, (i+1)*incy) += alpha * *at(xs, (i+1)*incx)
		*at(ys, (i+2)*incy) += alpha * *at(xs, (i+2)*incx)
		*at(ys, (i+3)*incy) += alpha * *at(xs, (i+3)*incx)
		i += 4
	}
	for ; n > 0; n-- {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
		i++
	}
}

//go:noinline
func AxpyUnsafeInlineR8(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	i := uintptr(0)
	n8 := n & mask8
	for ; i < n8; i += 8 {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
		*at(ys, (i+1)*incy) += alpha * *at(xs, (i+1)*incx)
		*at(ys, (i+2)*incy) += alpha * *at(xs, (i+2)*incx)
		*at(ys, (i+3)*incy) += alpha * *at(xs, (i+3)*incx)
		*at(ys, (i+4)*incy) += alpha * *at(xs, (i+4)*incx)
		*at(ys, (i+5)*incy) += alpha * *at(xs, (i+5)*incx)
		*at(ys, (i+6)*incy) += alpha * *at(xs, (i+6)*incx)
		*at(ys, (i+7)*incy) += alpha * *at(xs, (i+7)*incx)
	}
	for ; i < n; i++ {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
	}
}

//go:noinline
func AxpyUnsafeInlineXR8(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	i := uintptr(0)
	for ; n >= 8; n -= 8 {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
		*at(ys, (i+1)*incy) += alpha * *at(xs, (i+1)*incx)
		*at(ys, (i+2)*incy) += alpha * *at(xs, (i+2)*incx)
		*at(ys, (i+3)*incy) += alpha * *at(xs, (i+3)*incx)
		*at(ys, (i+4)*incy) += alpha * *at(xs, (i+4)*incx)
		*at(ys, (i+5)*incy) += alpha * *at(xs, (i+5)*incx)
		*at(ys, (i+6)*incy) += alpha * *at(xs, (i+6)*incx)
		*at(ys, (i+7)*incy) += alpha * *at(xs, (i+7)*incx)
		i += 8
	}
	for ; n > 0; n-- {
		*at(ys, (i+0)*incy) += alpha * *at(xs, (i+0)*incx)
		i++
	}
}

//go:noinline
func AxpyPointerR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	const Size = unsafe.Sizeof(xs[0])

	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))

	xn4 := unsafe.Add(xp, (n&mask4)*incx*Size)
	xn := unsafe.Add(xp, n*incx*Size)
	for uintptr(xp) < uintptr(xn4) {
		*(*float32)(unsafe.Add(yp, 0*Size)) += alpha * *(*float32)(unsafe.Add(xp, 0*Size))
		*(*float32)(unsafe.Add(yp, 1*Size)) += alpha * *(*float32)(unsafe.Add(xp, 1*Size))
		*(*float32)(unsafe.Add(yp, 2*Size)) += alpha * *(*float32)(unsafe.Add(xp, 2*Size))
		*(*float32)(unsafe.Add(yp, 3*Size)) += alpha * *(*float32)(unsafe.Add(xp, 3*Size))
		xp, yp = unsafe.Add(xp, 4*incx*Size), unsafe.Add(yp, 4*incy*Size)
	}
	for uintptr(xp) < uintptr(xn) {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, incx*Size), unsafe.Add(yp, incy*Size)
	}
}

//go:noinline
func AxpyPointerLoopR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	const Size = unsafe.Sizeof(xs[0])

	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))

	i := uintptr(0)
	n4 := n & mask4
	for ; i < n4; i += 4 {
		*(*float32)(unsafe.Add(yp, 0*Size)) += alpha * *(*float32)(unsafe.Add(xp, 0*Size))
		*(*float32)(unsafe.Add(yp, 1*Size)) += alpha * *(*float32)(unsafe.Add(xp, 1*Size))
		*(*float32)(unsafe.Add(yp, 2*Size)) += alpha * *(*float32)(unsafe.Add(xp, 2*Size))
		*(*float32)(unsafe.Add(yp, 3*Size)) += alpha * *(*float32)(unsafe.Add(xp, 3*Size))
		xp, yp = unsafe.Add(xp, 4*incx*Size), unsafe.Add(yp, 4*incy*Size)
	}
	for ; i < n; i++ {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, incx*Size), unsafe.Add(yp, incy*Size)
	}
}

//go:noinline
func AxpyPointerLoopXR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	const Size = unsafe.Sizeof(xs[0])

	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))

	for ; n >= 4; n -= 4 {
		*(*float32)(unsafe.Add(yp, 0*Size)) += alpha * *(*float32)(unsafe.Add(xp, 0*Size))
		*(*float32)(unsafe.Add(yp, 1*Size)) += alpha * *(*float32)(unsafe.Add(xp, 1*Size))
		*(*float32)(unsafe.Add(yp, 2*Size)) += alpha * *(*float32)(unsafe.Add(xp, 2*Size))
		*(*float32)(unsafe.Add(yp, 3*Size)) += alpha * *(*float32)(unsafe.Add(xp, 3*Size))
		xp, yp = unsafe.Add(xp, 4*incx*Size), unsafe.Add(yp, 4*incy*Size)
	}
	for ; n > 0; n-- {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, incx*Size), unsafe.Add(yp, incy*Size)
	}
}

//go:noinline
func AxpyPointerLoopInterleaveR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	const Size = unsafe.Sizeof(xs[0])

	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))

	i := uintptr(0)
	n4 := n & mask4
	for ; i < n4; i += 4 {
		x0 := *(*float32)(unsafe.Add(xp, 0*Size))
		x1 := *(*float32)(unsafe.Add(xp, 1*Size))
		x2 := *(*float32)(unsafe.Add(xp, 2*Size))
		x3 := *(*float32)(unsafe.Add(xp, 3*Size))

		m0 := alpha * x0
		m1 := alpha * x1
		m2 := alpha * x2
		m3 := alpha * x3

		t0 := *(*float32)(unsafe.Add(yp, 0*Size)) + m0
		t1 := *(*float32)(unsafe.Add(yp, 1*Size)) + m1
		t2 := *(*float32)(unsafe.Add(yp, 2*Size)) + m2
		t3 := *(*float32)(unsafe.Add(yp, 3*Size)) + m3

		*(*float32)(unsafe.Add(yp, 0*Size)) = t0
		*(*float32)(unsafe.Add(yp, 1*Size)) = t1
		*(*float32)(unsafe.Add(yp, 2*Size)) = t2
		*(*float32)(unsafe.Add(yp, 3*Size)) = t3

		xp, yp = unsafe.Add(xp, 4*incx*Size), unsafe.Add(yp, 4*incy*Size)
	}
	for ; i < n; i++ {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, incx*Size), unsafe.Add(yp, incy*Size)
	}
}

//go:noinline
func AxpyPointerLoopInterleaveXR4(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	const Size = unsafe.Sizeof(xs[0])

	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))

	for ; n >= 4; n -= 4 {
		x0 := *(*float32)(unsafe.Add(xp, 0*Size))
		x1 := *(*float32)(unsafe.Add(xp, 1*Size))
		x2 := *(*float32)(unsafe.Add(xp, 2*Size))
		x3 := *(*float32)(unsafe.Add(xp, 3*Size))

		m0 := alpha * x0
		m1 := alpha * x1
		m2 := alpha * x2
		m3 := alpha * x3

		t0 := *(*float32)(unsafe.Add(yp, 0*Size)) + m0
		t1 := *(*float32)(unsafe.Add(yp, 1*Size)) + m1
		t2 := *(*float32)(unsafe.Add(yp, 2*Size)) + m2
		t3 := *(*float32)(unsafe.Add(yp, 3*Size)) + m3

		*(*float32)(unsafe.Add(yp, 0*Size)) = t0
		*(*float32)(unsafe.Add(yp, 1*Size)) = t1
		*(*float32)(unsafe.Add(yp, 2*Size)) = t2
		*(*float32)(unsafe.Add(yp, 3*Size)) = t3

		xp, yp = unsafe.Add(xp, 4*incx*Size), unsafe.Add(yp, 4*incy*Size)
	}
	for ; n > 0; n-- {
		*(*float32)(yp) += alpha * *(*float32)(xp)
		xp, yp = unsafe.Add(xp, incx*Size), unsafe.Add(yp, incy*Size)
	}
}

//go:noinline
func AxpyPointerR4Alt(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr) {
	const Size = unsafe.Sizeof(xs[0])

	xp, yp := unsafe.SliceData(xs), unsafe.SliceData(ys)
	xn := offset(xp, n*incx)
	xn4 := offset(xp, (n&mask4)*incx)

	for less(xp, xn4) {
		*offset(yp, 0) += alpha * *offset(xp, 0)
		*offset(yp, 1) += alpha * *offset(xp, 1)
		*offset(yp, 2) += alpha * *offset(xp, 2)
		*offset(yp, 3) += alpha * *offset(xp, 3)
		xp, yp = offset(xp, 4*incx), offset(yp, 4*incy)
	}
	for less(xp, xn) {
		*yp += alpha * *xp
		xp, yp = offset(xp, incx), offset(yp, incy)
	}
}

func offset[T any](x *T, count uintptr) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(x), count*unsafe.Sizeof(*x)))
}
func less[T any](x, y *T) bool {
	return uintptr(unsafe.Pointer(x)) < uintptr(unsafe.Pointer(y))
}
