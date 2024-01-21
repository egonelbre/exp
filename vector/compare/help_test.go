package compare

func equalFloats(xs, ys []float32) bool {
	const epsilon = 1e-3
	
	if len(xs) != len(ys) {
		return false
	}

	for i, x := range xs {
		delta := ys[i] - x
		if delta < 0 {
			delta = -delta
		}
		if delta > epsilon {
			return false
		}
	}

	return true
}