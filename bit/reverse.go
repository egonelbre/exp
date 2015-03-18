package bit

func ReverseBits(x, width uint64) (rx uint64) {
	switch width {
	case 1, 2, 3, 4, 5, 6, 7:
		return reverseByte[x] >> (8 - width)
	case 8:
		return reverseByte[byte(x)]
	case 16:
		return reverseByte[byte(x)]<<8 | reverseByte[byte(x>>8)]
	case 24:
		return reverseByte[byte(x)]<<16 | reverseByte[byte(x>>16)] | reverseByte[byte(x>>8)]<<8
	case 32:
		return reverseByte[byte(x)]<<24 | reverseByte[byte(x>>8)]<<16 | reverseByte[byte(x>>16)]<<8 | reverseByte[byte(x>>24)]
	}
	return slowReverseBits(x, width)
}

// slow version
func slowReverseBits(x, width uint64) (rx uint64) {
	for i := uint64(0); i < width; i += 1 {
		rx = (rx << 1) | x&1
		x >>= 1
	}
	return
}
