package qpsortint

import (
	"testing"
)

func TestSort(t *testing.T) {
	var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	Sort(data)

	if !IsSorted(data) {
		t.Errorf("   got %v", data)
	}
}

func TestSortBasic(t *testing.T) {
	var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	SortBasic(data)

	if !IsSorted(data) {
		t.Errorf("   got %v", data)
	}
}
