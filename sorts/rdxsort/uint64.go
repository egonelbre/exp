package rdxsort

func Uint64(src, buf []uint64) {
	if len(src) != len(buf) {
		panic("len(src) != len(buf)")
	}

	var count [8][256]uint32
	for _, v := range src {
		count[0][(v>>(0*8))&0xFF]++
		count[1][(v>>(1*8))&0xFF]++
		count[2][(v>>(2*8))&0xFF]++
		count[3][(v>>(3*8))&0xFF]++
		count[4][(v>>(4*8))&0xFF]++
		count[5][(v>>(5*8))&0xFF]++
		count[6][(v>>(6*8))&0xFF]++
		count[7][(v>>(7*8))&0xFF]++
	}

	var offset [8][256]uint32
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
		offset[4][k] = offset[4][k-1] + count[4][k-1]
		offset[5][k] = offset[5][k-1] + count[5][k-1]
		offset[6][k] = offset[6][k-1] + count[6][k-1]
		offset[7][k] = offset[7][k-1] + count[7][k-1]
	}

	dst := buf
	for k := uint8(0); k < 8; k++ {
		off := &offset[k]

		p := k * 8
		for _, v := range src {
			key := uint8(v >> p)
			dst[off[key]] = v
			off[key]++
		}

		dst, src = src, dst
	}
}

func Uint32(src, buf []uint32) {
	if len(src) != len(buf) {
		panic("len(src) != len(buf)")
	}

	var count [4][256]uint32
	for _, v := range src {
		count[0][(v>>(0*8))&0xFF]++
		count[1][(v>>(1*8))&0xFF]++
		count[2][(v>>(2*8))&0xFF]++
		count[3][(v>>(3*8))&0xFF]++
	}

	var offset [4][256]uint32
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
	}

	dst := buf
	for k := uint8(0); k < 4; k++ {
		off := &offset[k]

		p := k * 8
		for _, v := range src {
			key := uint8(v >> p)
			dst[off[key]] = v
			off[key]++
		}

		dst, src = src, dst
	}
}
