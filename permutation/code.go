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
