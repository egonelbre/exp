package compare

// AsmAxpyInc is
//
//	for i := 0; i < int(n); i++ {
//		y[iy] += alpha * x[ix]
//		ix += incX
//		iy += incY
//	}
func AsmAxpyInc(alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr)
