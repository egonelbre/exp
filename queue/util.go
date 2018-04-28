package queue

import "math/bits"

func ceil(a, n int) int {
	r := ((a + n - 1) / n) * n
	if r <= n {
		return n * 2
	}
	return r
}

func nextPowerOfTwo(v uint32) uint32 {
	return 1 << uint(32-bits.LeadingZeros32(v))
}
