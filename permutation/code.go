package permutation

const base = 12

var fact [base]int

func init() {
	fact[0] = 1
	for i := 1; i < base; i++ {
		fact[i] = fact[i-1] * i
	}
	for i := 0; i < base/2; i++ {
		fact[i], fact[base-i-1] = fact[base-i-1], fact[i]
	}
}

func Code(perm [base]byte) int { return CodeCopy(perm) }

func CodeCopy(perm [base]byte) int {
	r := 0
	for min := byte(0); min < base-1; min++ {
		z := 0
		n := base - min
		for i, v := range perm[:n] {
			if v == min {
				z = i
				break
			}
		}
		copy(perm[z:], perm[z+1:n])
		r += z * fact[min]
	}
	return r
}

func CodeCount(perm [base]byte) int {
	r := 0
	for min := byte(0); min < base-1; min++ {
		z := 0
		for _, v := range perm {
			if v > min {
				z++
			} else if v == min {
				break
			}
		}
		r += z * fact[min]
	}
	return r
}

func CodeTable(perm [base]byte) int {
	var index [base]byte
	for i, v := range perm {
		index[v] = byte(i)
	}

	r := 0
	for min := byte(0); min < base-1; min++ {
		z := 0
		for _, v := range perm[:index[min]] {
			if v > min {
				z++
			}
		}
		r += z * fact[min]
	}
	return r
}

func CodeTable2(perm [base]byte) int {
	var index [base]byte
	for i, v := range perm {
		index[v] = byte(i)
	}

	r := 0
	for min := byte(0); min < base-1; min++ {
		z := index[min]
		for i, pos := range index {
			if pos > z {
				index[i] = pos - 1
			}
		}
		r += int(z) * fact[min]
	}
	return r
}

func PackNybble(perm [base]byte) int {
	r := 0
	for _, v := range perm {
		r = r<<4 + int(v)
	}
	return r
}

func PackTight(perm [base]byte) int {
	r := 0
	for _, v := range perm {
		r = r*12 + int(v)
	}
	return r
}

func PackNybbleUnroll(perm [base]byte) int {
	return int(perm[0]) +
		int(perm[1])<<(4*1) +
		int(perm[2])<<(4*2) +
		int(perm[3])<<(4*3) +
		int(perm[4])<<(4*4) +
		int(perm[5])<<(4*5) +
		int(perm[6])<<(4*6) +
		int(perm[7])<<(4*7) +
		int(perm[8])<<(4*8) +
		int(perm[9])<<(4*9) +
		int(perm[10])<<(4*10) +
		int(perm[11])<<(4*11)
}

func PackNybbleUnroll2(perm [base]byte) int {
	i0 := int(perm[0]) | int(perm[1])<<8 | int(perm[2])<<16 | int(perm[3])<<24
	i1 := int(perm[4]) | int(perm[5])<<8 | int(perm[6])<<16 | int(perm[7])<<24
	i2 := int(perm[8]) | int(perm[9])<<8 | int(perm[10])<<16 | int(perm[11])<<24
	return i0 | i1<<4 | i2<<32
}

const (
	b12_0  = 1
	b12_1  = b12_0 * 12
	b12_2  = b12_1 * 12
	b12_3  = b12_2 * 12
	b12_4  = b12_3 * 12
	b12_5  = b12_4 * 12
	b12_6  = b12_5 * 12
	b12_7  = b12_6 * 12
	b12_8  = b12_7 * 12
	b12_9  = b12_8 * 12
	b12_10 = b12_9 * 12
	b12_11 = b12_10 * 12
)

func PackTightUnroll(perm [base]byte) int {
	return int(perm[0])*b12_0 +
		int(perm[1])*b12_1 +
		int(perm[2])*b12_2 +
		int(perm[3])*b12_3 +
		int(perm[4])*b12_4 +
		int(perm[5])*b12_5 +
		int(perm[6])*b12_6 +
		int(perm[7])*b12_7 +
		int(perm[8])*b12_8 +
		int(perm[9])*b12_9 +
		int(perm[10])*b12_10 +
		int(perm[11])*b12_11
}
