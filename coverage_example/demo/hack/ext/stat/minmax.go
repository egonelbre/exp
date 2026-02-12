package stat

func MinMax(values []int) (minv, maxv int) {
	if len(values) == 0 {
		return 0, 0
	}
	minv = values[0]
	maxv = values[0]
	for _, v := range values {
		minv = min(minv, v)
		maxv = max(maxv, v)
	}
	return
}
