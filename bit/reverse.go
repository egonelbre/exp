package bit

func ReverseBits(x, width uint) (rx uint) {
	for i := uint(0); i < width; i += 1 {
		rx = (rx << 1) | x&1
		x >>= 1
	}
	return
}
