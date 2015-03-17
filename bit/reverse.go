package bit

func ReverseBits(x, width uint64) (rx uint64) {
	for i := uint64(0); i < width; i += 1 {
		rx = (rx << 1) | x&1
		x >>= 1
	}
	return
}
