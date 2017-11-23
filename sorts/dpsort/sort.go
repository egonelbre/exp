//
// based on http://codeblab.com/wp-content/uploads/2009/09/DualPivotQuicksort.pdf
//

package dpsort

import "sort"

const insertionSortThreshold = 27

func dataequal(data sort.Interface, i, k int) bool {
    return !(data.Less(i, k) || data.Less(k, i))
}

func insertionSort(data sort.Interface, lo, hi int) {
    for i := lo + 1; i <= hi; i++ {
        for k := i; k > lo && data.Less(k, k-1); k-- {
            data.Swap(k, k-1)
        }
    }
}

func Sort(data sort.Interface) {
    sortp(data, 0, data.Len()-1, 3)
}

func sortp(data sort.Interface, left, right int, div int) {
    n := right - left
    if n == 0 {
        return
    } else if n < insertionSortThreshold {
        insertionSort(data, left, right)
        return
    }

    third := n / div
    // "medians"
    m1 := left + third
    m2 := right - third
    if m1 <= left {
        m1 = left + 1
    }
    if m2 >= right {
        m2 = right - 1
    }

    if data.Less(m1, m2) {
        data.Swap(m1, left)
        data.Swap(m2, right)
    } else {
        data.Swap(m1, right)
        data.Swap(m2, left)
    }

    // pointers
    less := left + 1
    great := right - 1
    // sorting
    for k := less; k <= great; k++ {
        if data.Less(k, left) {
            data.Swap(k, less)
            less++
        } else if data.Less(right, k) {
            for k < great && data.Less(right, great) {
                great--
            }
            data.Swap(k, great)
            great--
            if data.Less(k, left) {
                data.Swap(k, less)
                less++
            }
        }
    }

    // swaps
    dist := great - less
    if dist < 13 {
        div++
    }

    data.Swap(less-1, left)
    data.Swap(great+1, right)

    // subarrays
    sortp(data, left, less-2, div)
    sortp(data, great+2, right, div)

    // equal elements
    if dist > n-13 && !dataequal(data, left, right) {
        for k := less; k <= great; k++ {
            if dataequal(data, k, left) {
                data.Swap(k, less)
                less++
            } else if dataequal(data, k, right) {
                data.Swap(k, great)
                great--
                if dataequal(data, k, left) {
                    data.Swap(k, less)
                    less++
                }
            }
        }
    }
    // subarray
    if data.Less(left, right) {
        sortp(data, less, great, div)
    }
}
