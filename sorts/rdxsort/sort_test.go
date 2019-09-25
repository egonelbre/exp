package rdxsort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/egonelbre/exp/sorts/rdxsort"
)

func TestSort(t *testing.T) {
	var data = []uint64{2, 1}
	buf := make([]uint64, len(data))
	rdxsort.Uint64(data, buf)
	if !isSorted(data) {
		t.Errorf("   got %v", data)
	}
}

func TestRandom(t *testing.T) {
	for _, size := range []int{1, 2, 3, 4, 10, 32, 64, 100, 1000, 10000} {
		for k := 0; k < 10; k++ {
			data := make([]uint64, size)
			for i := range data {
				data[i] = rand.Uint64()
			}

			rdxsort.Uint64(data, make([]uint64, size))
			if !isSorted(data) {
				t.Errorf("   got %v", data)
			}
		}
	}
}

func isSorted(vs []uint64) bool {
	return sort.SliceIsSorted(vs, func(i, k int) bool { return vs[i] < vs[k] })
}
