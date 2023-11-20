package vecprocess

// AxpyInc is
//
//	for i := 0; i < int(n); i++ {
//		y[iy] += alpha * x[ix]
//		ix += incX
//		iy += incY
//	}
func AxpyInc(alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr)
