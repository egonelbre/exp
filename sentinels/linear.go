package sentinels

import "unsafe"

func SearchBasic(items []int, elem int) int {
	for i, v := range items {
		if v == elem {
			return i
		}
	}
	return len(items)
}

// requires -gcflags="-B" to be effective
func SearchBreak(items []int, elem int) (i int) {
	for ; i < len(items); i++ {
		if items[i] == elem {
			break
		}
	}
	return i
}

// requires -gcflags="-B" to be effective
func SearchSentinel(items []int, elem int) (i int) {
	var last int
	last, items[len(items)-1] = items[len(items)-1], elem
	for items[i] != elem {
		i++
	}
	if i+1 == len(items) && last != elem {
		i++
	}
	items[len(items)-1] = last
	return i
}

func SearchUnsafe(items []int, elem int) (i int) {
	var last int
	last, items[len(items)-1] = items[len(items)-1], elem

	p := unsafe.Pointer(&items[0])
	for *(*int)(p) != elem {
		p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(elem))
	}
	i = int((uintptr(p) - uintptr(unsafe.Pointer(&items[0]))) / unsafe.Sizeof(elem))
	if i+1 == len(items) && last != elem {
		i++
	}
	items[len(items)-1] = last
	return i
}
