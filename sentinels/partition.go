package sentinels

// requires -gcflags="-B" to be effective
func PartitionBreak(items []int, pivot int) int {
	items[pivot], items[0] = items[0], items[pivot]
	lo, hi := 1, len(items)-1
loop:
	for ; ; lo, hi = lo+1, hi-1 {
		for ; ; lo++ {
			if lo > hi {
				break loop
			}
			if items[lo] >= items[0] {
				break
			}
		}
		for ; ; hi-- {
			if lo >= hi {
				break loop
			}
			if items[0] >= items[hi] {
				break
			}
		}
		items[hi], items[lo] = items[lo], items[hi]
	}
	lo--
	items[lo], items[0] = items[0], items[lo]
	return lo
}

// requires -gcflags="-B" to be effective
func PartitionSentinel(items []int, pivot int) int {
	if len(items) <= 1 {
		return 0
	}
	items[pivot], items[0] = items[0], items[pivot]

	pv, save := items[0], items[0]
	save, items[len(items)-1] = items[len(items)-1], save

	lo, hi := 0, len(items)-1
	for {
		for {
			lo++
			if items[lo] < pv {
				break
			}
		}
		items[hi] = items[lo]
		for {
			hi--
			if items[hi] >= pv {
				break
			}
		}
		if lo >= hi {
			break
		}
		items[lo] = items[hi]
	}

	if lo == hi+2 {
		items[lo] = items[hi+1]
		lo--
	}
	items[lo] = save
	if pv < save {
		lo--
	}
	items[lo], items[0] = items[0], items[lo]
	return lo
}
