package bit

// AbsEncode converts number from signed to unsigned representation
// The sign will be encoded as the least significant bit in the result
// AbsEncode(x) < abs(x)*2
func AbsEncode(x int64) uint64 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return uint64(-x) << 1
	}
	return uint64(x<<1) - 1
}

// AbsDecode is the inverse operation of AbsEncode
// AbsDecode(AbsEncode(x)) == x
func AbsDecode(x uint64) int64 {
	if x == 0 {
		return 0
	}
	if x%2 == 0 {
		return -int64(x >> 1)
	}
	return int64((x + 1) >> 1)
}
