package rdxsort

func Uint64(src, buf []uint64) {
	if int64(len(src)) >= int64(1<<32) {
		panic("slice too large")
	}
	if len(src) != len(buf) {
		panic("len(src) != len(buf)")
	}

	var count [8][256]uint32
	for _, v := range src {
		count[0][byte(v>>(0*8))]++
		count[1][byte(v>>(1*8))]++
		count[2][byte(v>>(2*8))]++
		count[3][byte(v>>(3*8))]++
		count[4][byte(v>>(4*8))]++
		count[5][byte(v>>(5*8))]++
		count[6][byte(v>>(6*8))]++
		count[7][byte(v>>(7*8))]++
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

	var nonZero [8]byte
	for k := range nonZero {
		nz := byte(0)
		for _, v := range count[k] {
			if v != 0 {
				nz++
			}
		}
		nonZero[k] = nz
	}

	dst := buf
	for k := uint8(0); k < 8; k++ {
		if nonZero[k] == 1 {
			continue
		}
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
	if int64(len(src)) >= int64(1<<32) {
		panic("slice too large")
	}
	if len(src) != len(buf) {
		panic("len(src) != len(buf)")
	}

	var count [4][256]uint32
	for _, v := range src {
		count[0][byte(v>>(0*8))]++
		count[1][byte(v>>(1*8))]++
		count[2][byte(v>>(2*8))]++
		count[3][byte(v>>(3*8))]++
	}

	var offset [4][256]uint32
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
	}

	var nonZero [4]byte
	for k := range nonZero {
		nz := byte(0)
		for _, v := range count[k] {
			if v != 0 {
				nz++
			}
		}
		nonZero[k] = nz
	}

	dst := buf
	for k := uint8(0); k < 4; k++ {
		if nonZero[k] == 1 {
			continue
		}
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
