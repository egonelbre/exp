/*

Based on go std sort

*/

package stdsortint

func dataless(data []int, i, k int) bool { return data[i] < data[k] }
func dataswap(data []int, i, k int)      { data[i], data[k] = data[k], data[i] }

// Insertion sort
func insertionSort(data []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && dataless(data, j, j-1); j-- {
			dataswap(data, j, j-1)
		}
	}
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && dataless(data, first+child, first+child+1) {
			child++
		}
		if !dataless(data, first+root, first+child) {
			return
		}
		dataswap(data, first+root, first+child)
		root = child
	}
}

func heapSort(data []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		dataswap(data, first, first+i)
		siftDown(data, lo, i, first)
	}
}

// Quicksort, loosely following Bentley and McIlroy,
// ``Engineering a Sort Function,'' SP&E November 1993.

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(data []int, m1, m0, m2 int) {
	// sort 3 elements
	if dataless(data, m1, m0) {
		dataswap(data, m1, m0)
	}
	// data[m0] <= data[m1]
	if dataless(data, m2, m1) {
		dataswap(data, m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if dataless(data, m1, m0) {
			dataswap(data, m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func swapRange(data []int, a, b, n int) {
	for i := 0; i < n; i++ {
		dataswap(data, a+i, b+i)
	}
}

func doPivot(data []int, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && dataless(data, a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !dataless(data, pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && dataless(data, pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		dataswap(data, b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !dataless(data, pivot, hi-1) { // data[hi-1] = pivot
			dataswap(data, c, hi-1)
			c++
			dups++
		}
		if !dataless(data, b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !dataless(data, m, pivot) { // data[m] = pivot
			dataswap(data, m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !dataless(data, b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && dataless(data, a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			dataswap(data, a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	dataswap(data, pivot, b-1)
	return b - 1, c
}

func quickSort(data []int, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(data, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if dataless(data, i, i-6) {
				dataswap(data, i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// dataless data, and dataswap.data,  The sort is not guaranteed to be stable.
func Sort(data []int) {
	n := len(data)
	quickSort(data, 0, n, maxDepth(n))
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

// IsSorted reports whether data is sorted.
func IsSorted(data []int) bool {
	n := len(data)
	for i := n - 1; i > 0; i-- {
		if dataless(data, i, i-1) {
			return false
		}
	}
	return true
}
