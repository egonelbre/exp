package qpsort

import "sort"

const insertionSortThreshold = 48

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

func insertionSort(data sort.Interface, lo, hi int) {
	for i := lo + 1; i <= hi; i++ {
		for k := i; k > lo && data.Less(k, k-1); k-- {
			data.Swap(k, k-1)
		}
	}
}

func SortBasic(data sort.Interface) {
	sort3basic(data, 0, data.Len()-1)
}

func Sort(data sort.Interface) {
	sort3(data, 0, data.Len()-1)
}

func sort3basic(data sort.Interface, lo, hi int) {
	if hi == lo {
		return
	} else if hi-lo < insertionSortThreshold {
		insertionSort(data, lo, hi)
		return
	}

	m := (lo + hi) >> 1 // TOFIX: potential overflow
	if data.Less(m, lo) {
		data.Swap(m, lo)
	}

	if data.Less(hi, m) {
		data.Swap(hi, m)
		if data.Less(m, lo) {
			data.Swap(m, lo)
		}
	}

	data.Swap(lo+1, m)

	// Pointers a and b initially point to the first element of the array while c
	// and d initially point to the last element of the array.
	a := lo + 2
	b := lo + 2
	c := hi - 1
	d := hi - 1

	for b <= c {
		for data.Less(b, lo+1) && b <= c {
			if data.Less(b, lo) {
				data.Swap(a, b)
				a++
			}
			b++
		}

		for data.Less(lo+1, c) && b <= c {
			if data.Less(hi, c) {
				data.Swap(c, d)
				d--
			}
			c--
		}

		if b <= c {
			if data.Less(hi, b) {
				if data.Less(c, lo) {
					data.Swap(b, a)
					data.Swap(a, c)
					a++
				} else {
					data.Swap(b, c)
				}
				data.Swap(c, d)
				b++
				c--
				d--
			} else {
				if data.Less(c, lo) {
					data.Swap(b, a)
					data.Swap(a, c)
					a++
				} else {
					data.Swap(b, c)
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
	data.Swap(lo+1, a)
	data.Swap(a, b)

	a--
	data.Swap(lo, a)
	data.Swap(hi, d)

	sort3basic(data, lo, a-1)
	sort3basic(data, a+1, b)
	sort3basic(data, c, d-1)
	sort3basic(data, d+1, hi)
}

func sort3(data sort.Interface, lo, hi int) {
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
		if data.Less(e2, e1) {
			data.Swap(e2, e1)
		}

		if data.Less(e3, e2) {
			data.Swap(e3, e2)
			if data.Less(e2, e1) {
				data.Swap(e2, e1)
			}
		}

		if data.Less(e4, e3) {
			data.Swap(e4, e3)
			if data.Less(e3, e2) {
				data.Swap(e3, e2)
				if data.Less(e2, e1) {
					data.Swap(e2, e1)
				}
			}
		}

		if data.Less(e5, e4) {
			data.Swap(e5, e4)
			if data.Less(e4, e3) {
				data.Swap(e4, e3)
				if data.Less(e3, e2) {
					data.Swap(e3, e2)
					if data.Less(e2, e1) {
						data.Swap(e2, e1)
					}
				}
			}
		}

		if data.Less(e6, e5) {
			data.Swap(e6, e5)
			if data.Less(e5, e4) {
				data.Swap(e5, e4)
				if data.Less(e4, e3) {
					data.Swap(e4, e3)
					if data.Less(e3, e2) {
						data.Swap(e3, e2)
						if data.Less(e2, e1) {
							data.Swap(e2, e1)
						}
					}
				}
			}
		}

		if data.Less(e7, e6) {
			data.Swap(e7, e6)
			if data.Less(e6, e5) {
				data.Swap(e6, e5)
				if data.Less(e5, e4) {
					data.Swap(e5, e4)
					if data.Less(e4, e3) {
						data.Swap(e4, e3)
						if data.Less(e3, e2) {
							data.Swap(e3, e2)
							if data.Less(e2, e1) {
								data.Swap(e2, e1)
							}
						}
					}
				}
			}
		}
	}

	data.Swap(lo, e2)
	data.Swap(lo+1, e4)
	data.Swap(hi, e6)

	// Pointers a and b initially point to the first element of the array while c
	// and d initially point to the last element of the array.
	a := lo + 2
	b := lo + 2
	c := hi - 1
	d := hi - 1

	for b <= c {
		for data.Less(b, lo+1) && b <= c {
			if data.Less(b, lo) {
				data.Swap(a, b)
				a++
			}
			b++
		}

		for data.Less(lo+1, c) && b <= c {
			if data.Less(hi, c) {
				data.Swap(c, d)
				d--
			}
			c--
		}

		if b <= c {
			if data.Less(hi, b) {
				if data.Less(c, lo) {
					data.Swap(b, a)
					data.Swap(a, c)
					a++
				} else {
					data.Swap(b, c)
				}
				data.Swap(c, d)
				b++
				c--
				d--
			} else {
				if data.Less(c, lo) {
					data.Swap(b, a)
					data.Swap(a, c)
					a++
				} else {
					data.Swap(b, c)
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
	data.Swap(lo+1, a)
	data.Swap(a, b)
	a--
	data.Swap(lo, a)
	data.Swap(hi, d)

	sort3(data, lo, a-1)
	sort3(data, a+1, b)
	sort3(data, c, d-1)
	sort3(data, d+1, hi)
}
