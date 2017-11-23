/*

Based on https://github.com/dmcmanam/quicksort/blob/master/src/Quicksort.java

*/

package qpsortint

const insertionSortThreshold = 48

func dataless(data []int, i, k int) bool { return data[i] < data[k] }
func dataswap(data []int, i, k int)      { data[i], data[k] = data[k], data[i] }

func insertionSortLocal(data []int, lo, hi int) {
	if lo == hi {
		return
	}

	for i := lo + 1; i < hi+1; i++ {
		v := data[i]
		k := i
		for k > 0 && data[k-1] > v {
			data[k] = data[k-1]
			k--
		}
		data[k] = v
	}
}

func insertionSort(data []int, lo, hi int) {
	for i := lo + 1; i <= hi; i++ {
		for k := i; k > lo && dataless(data, k, k-1); k-- {
			dataswap(data, k, k-1)
		}
	}
}

func SortBasic(data []int) {
	sort3basic(data, 0, len(data)-1)
}

func Sort(data []int) {
	sort3(data, 0, len(data)-1)
}

func sort3basic(data []int, lo, hi int) {
	if hi == lo {
		return
	} else if hi-lo < insertionSortThreshold {
		insertionSort(data, lo, hi)
		return
	}

	m := (lo + hi) >> 1 // TOFIX: potential overflow
	if dataless(data, m, lo) {
		dataswap(data, m, lo)
	}

	if dataless(data, hi, m) {
		dataswap(data, hi, m)
		if dataless(data, m, lo) {
			dataswap(data, m, lo)
		}
	}

	dataswap(data, lo+1, m)

	// Pointers a and b initially point to the first element of the array while c
	// and d initially point to the last element of the array.
	a := lo + 2
	b := lo + 2
	c := hi - 1
	d := hi - 1

	for b <= c {
		for dataless(data, b, lo+1) && b <= c {
			if dataless(data, b, lo) {
				dataswap(data, a, b)
				a++
			}
			b++
		}

		for dataless(data, lo+1, c) && b <= c {
			if dataless(data, hi, c) {
				dataswap(data, c, d)
				d--
			}
			c--
		}

		if b <= c {
			if dataless(data, hi, b) {
				if dataless(data, c, lo) {
					dataswap(data, b, a)
					dataswap(data, a, c)
					a++
				} else {
					dataswap(data, b, c)
				}
				dataswap(data, c, d)
				b++
				c--
				d--
			} else {
				if dataless(data, c, lo) {
					dataswap(data, b, a)
					dataswap(data, a, c)
					a++
				} else {
					dataswap(data, b, c)
				}
				b++
				c--
			}
		}
	}

	a--
	b--
	c++
	d++
	dataswap(data, lo+1, a)
	dataswap(data, a, b)

	a--
	dataswap(data, lo, a)
	dataswap(data, hi, d)

	sort3basic(data, lo, a-1)
	sort3basic(data, a+1, b)
	sort3basic(data, c, d-1)
	sort3basic(data, d+1, hi)
}

func sort3(data []int, lo, hi int) {
	if lo == hi {
		return
	}

	length := hi - lo + 1
	if length < insertionSortThreshold {
		insertionSort(data, lo, hi)
		return
	}

	seventh := length>>3 + length>>6 + 1
	e4 := (lo + hi) >> 1 // TODO: fix overflow
	e3 := e4 - seventh
	e2 := e3 - seventh
	e1 := lo
	e5 := e4 + seventh
	e6 := e5 + seventh
	e7 := hi

	{ // inline insertion sort
		if dataless(data, e2, e1) {
			dataswap(data, e2, e1)
		}

		if dataless(data, e3, e2) {
			dataswap(data, e3, e2)
			if dataless(data, e2, e1) {
				dataswap(data, e2, e1)
			}
		}

		if dataless(data, e4, e3) {
			dataswap(data, e4, e3)
			if dataless(data, e3, e2) {
				dataswap(data, e3, e2)
				if dataless(data, e2, e1) {
					dataswap(data, e2, e1)
				}
			}
		}

		if dataless(data, e5, e4) {
			dataswap(data, e5, e4)
			if dataless(data, e4, e3) {
				dataswap(data, e4, e3)
				if dataless(data, e3, e2) {
					dataswap(data, e3, e2)
					if dataless(data, e2, e1) {
						dataswap(data, e2, e1)
					}
				}
			}
		}

		if dataless(data, e6, e5) {
			dataswap(data, e6, e5)
			if dataless(data, e5, e4) {
				dataswap(data, e5, e4)
				if dataless(data, e4, e3) {
					dataswap(data, e4, e3)
					if dataless(data, e3, e2) {
						dataswap(data, e3, e2)
						if dataless(data, e2, e1) {
							dataswap(data, e2, e1)
						}
					}
				}
			}
		}

		if dataless(data, e7, e6) {
			dataswap(data, e7, e6)
			if dataless(data, e6, e5) {
				dataswap(data, e6, e5)
				if dataless(data, e5, e4) {
					dataswap(data, e5, e4)
					if dataless(data, e4, e3) {
						dataswap(data, e4, e3)
						if dataless(data, e3, e2) {
							dataswap(data, e3, e2)
							if dataless(data, e2, e1) {
								dataswap(data, e2, e1)
							}
						}
					}
				}
			}
		}
	}

	dataswap(data, lo, e2)
	dataswap(data, lo+1, e4)
	dataswap(data, hi, e6)

	// Pointers a and b initially point to the first element of the array while c
	// and d initially point to the last element of the array.
	a := lo + 2
	b := lo + 2
	c := hi - 1
	d := hi - 1

	for b <= c {
		for dataless(data, b, lo+1) && b <= c {
			if dataless(data, b, lo) {
				dataswap(data, a, b)
				a++
			}
			b++
		}

		for dataless(data, lo+1, c) && b <= c {
			if dataless(data, hi, c) {
				dataswap(data, c, d)
				d--
			}
			c--
		}

		if b <= c {
			if dataless(data, hi, b) {
				if dataless(data, c, lo) {
					dataswap(data, b, a)
					dataswap(data, a, c)
					a++
				} else {
					dataswap(data, b, c)
				}
				dataswap(data, c, d)
				b++
				c--
				d--
			} else {
				if dataless(data, c, lo) {
					dataswap(data, b, a)
					dataswap(data, a, c)
					a++
				} else {
					dataswap(data, b, c)
				}
				b++
				c--
			}
		}
	}

	a--
	b--
	c++
	d++
	dataswap(data, lo+1, a)
	dataswap(data, a, b)
	a--
	dataswap(data, lo, a)
	dataswap(data, hi, d)

	sort3(data, lo, a-1)
	sort3(data, a+1, b)
	sort3(data, c, d-1)
	sort3(data, d+1, hi)
}

// IsSorted reports whether data is sorted.
func IsSorted(data []int) bool {
	n := len(data)
	for i := 1; i < n; i++ {
		if dataless(data, i, i-1) {
			return false
		}
	}
	return true
}
