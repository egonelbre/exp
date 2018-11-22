package lcs

// Basic returns the length of the largest common subsequence
func Basic(a, b []int64) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	curr := make([]int, len(b))
	prev := make([]int, len(b))

	max := 0

	for ai, ax := range a {
		for bi, bx := range b {
			if ax != bx {
				curr[bi] = 0
			} else {
				if ai == 0 || bi == 0 {
					curr[bi] = 1
				} else {
					curr[bi] = 1 + prev[bi-1]
				}

				if max < curr[bi] {
					max = curr[bi]
				}
			}
		}
		curr, prev = prev, curr
	}

	return max
}

// Lift returns the length of the largest common subsequence
func Lift(a, b []int64) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	curr := make([]int, len(b))
	prev := make([]int, len(b))

	max := 0
	for ai, ax := range a {
		if b[ai] == ax {
			prev[ai] = 1
		}
	}

	for ai := 1; ai < len(a); ai++ {
		ax := a[ai]

		if ax == b[0] {
			curr[0] = 1
		} else {
			curr[0] = 0
		}

		for bi := 1; bi < len(b); bi++ {
			bx := b[bi]
			if ax != bx {
				curr[bi] = 0
			} else {
				curr[bi] = 1 + prev[bi-1]
				if max < curr[bi] {
					max = curr[bi]
				}
			}
		}
		curr, prev = prev, curr
	}

	return max
}

// Wave returns the length of the largest common subsequence
func Wave(a, b []int64) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	prev := make([]int, len(b))
	mid := make([]int, len(b))
	curr := make([]int, len(b))

	max := 0
	//         prev
	//        / mid
	//       / / curr
	//      / / /
	// 0 1 2 3 4
	// 0 1 2 3
	// 0 1 2
	// 0 1
	// 0

	for wave := 0; wave < len(b); wave++ {
		for i := 0; i < wave; i++ {
			ax, bx := a[wave], b[wave-i]
			if ax == bx {
				if i == 0 {
					curr[i] = 1
				} else {
					curr[i] = prev[i-1] + 1
					if curr[i] > max {
						max = curr[i]
					}
				}
			} else {
				curr[i] = 0
			}
		}
		prev, mid, curr = mid, curr, prev
	}

	//           prev
	//          / mid
	//         / / curr
	//        / / /
	//       3 4 .
	//     2 3 3
	//   1 2 2 2
	// 0 1 1 1 1
	// 0 0 0 0 0

	for wave := len(b) - 2; wave >= 0; wave-- {
		for i := 0; i < wave; i++ {
			ax, bx := a[wave], b[wave-i]
			if ax == bx {
				if i == 0 {
					curr[i] = 1
				} else {
					curr[i] = prev[i-1] + 1
					if curr[i] > max {
						max = curr[i]
					}
				}
			} else {
				curr[i] = 0
			}
		}
		prev, mid, curr = mid, curr, prev
	}

	return max
}
