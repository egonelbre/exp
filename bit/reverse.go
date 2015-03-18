package bit

func Reverse(x uint64, width uint) (rx uint64) {
	switch width {
	case 1, 2, 3, 4, 5, 6, 7:
		return uint64(reverseByte[x] >> (8 - width))
	case 8:
		return uint64(reverseByte[byte(x)])
	case 16:
		return uint64(reverseByte[byte(x)])<<8 | uint64(reverseByte[byte(x>>8)])
	case 24:
		return uint64(reverseByte[byte(x)])<<16 | uint64(reverseByte[byte(x>>16)]) | uint64(reverseByte[byte(x>>8)])<<8
	case 32:
		return uint64(reverseByte[byte(x)])<<24 | uint64(reverseByte[byte(x>>8)])<<16 | uint64(reverseByte[byte(x>>16)])<<8 | uint64(reverseByte[byte(x>>24)])
	}
	return slowReverse(x, width)
}

// slow version
func slowReverse(x uint64, width uint) (rx uint64) {
	for i := uint(0); i < width; i += 1 {
		rx = (rx << 1) | x&1
		x >>= 1
	}
	return
}
