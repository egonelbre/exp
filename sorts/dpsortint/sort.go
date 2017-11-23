//
// based on http://codeblab.com/wp-content/uploads/2009/09/DualPivotQuicksort.pdf
//

package dpsortint

const insertionSortThreshold = 27

func dataequal(data []int, i, k int) bool { return !(dataless(data, i, k) || dataless(data, k, i)) }
func dataless(data []int, i, k int) bool  { return data[i] < data[k] }
func dataswap(data []int, i, k int)       { data[i], data[k] = data[k], data[i] }

func insertionSort(data []int, lo, hi int) {
    for i := lo + 1; i <= hi; i++ {
        for k := i; k > lo && dataless(data, k, k-1); k-- {
            dataswap(data, k, k-1)
        }
    }
}

func Sort(data []int) {
    sortp(data, 0, len(data)-1, 3)
}

func sortp(data []int, left, right int, div int) {
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

    if dataless(data, m1, m2) {
        dataswap(data, m1, left)
        dataswap(data, m2, right)
    } else {
        dataswap(data, m1, right)
        dataswap(data, m2, left)
    }

    // pointers
    less := left + 1
    great := right - 1
    // sorting
    for k := less; k <= great; k++ {
        if dataless(data, k, left) {
            dataswap(data, k, less)
            less++
        } else if dataless(data, right, k) {
            for k < great && dataless(data, right, great) {
                great--
            }
            dataswap(data, k, great)
            great--
            if dataless(data, k, left) {
                dataswap(data, k, less)
                less++
            }
        }
    }

    // swaps
    dist := great - less
    if dist < 13 {
        div++
    }

    dataswap(data, less-1, left)
    dataswap(data, great+1, right)

    // subarrays
    sortp(data, left, less-2, div)
    sortp(data, great+2, right, div)

    // equal elements
    if dist > n-13 && !dataequal(data, left, right) {
        for k := less; k <= great; k++ {
            if dataequal(data, k, left) {
                dataswap(data, k, less)
                less++
            } else if dataequal(data, k, right) {
                dataswap(data, k, great)
                great--
                if dataequal(data, k, left) {
                    dataswap(data, k, less)
                    less++
                }
            }
        }
    }
    // subarray
    if dataless(data, left, right) {
        sortp(data, less, great, div)
    }
}
