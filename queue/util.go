package queue

func ceil(a, n int) int {
	r := ((a + n - 1) / n) * n
	if r <= n {
		return n * 2
	}
	return r
}
