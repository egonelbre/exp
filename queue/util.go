package queue

import (
	"math/bits"
	"runtime"
	"time"
)

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

func backoff(np *int) {
	n := *np
	*np++

	if n < 3 {
		return
	} else if n < 10 {
		runtime.Gosched()
	} else if n < 12 {
		time.Sleep(0) // osyield
	} else {
		time.Sleep(10 * time.Microsecond)
	}

}
