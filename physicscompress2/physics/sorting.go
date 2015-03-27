package physics

import (
	"sort"

	"github.com/egonelbre/exp/bit"
)

type sorterZ struct {
	order []int
	get   func(i int) int
}

func (s *sorterZ) Len() int      { return len(s.order) }
func (s *sorterZ) Swap(i, j int) { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *sorterZ) Less(i, j int) bool {
	vi := int64(s.get(s.order[i]))
	vj := int64(s.get(s.order[j]))
	return bit.ZEncode(vi) < bit.ZEncode(vj)
}

func SortByZ(items []int, get func(i int) int) {
	sort.Sort(&sorterZ{
		order: items,
		get:   get,
	})
}

type sorterZCount struct {
	order []int
	get   func(i int) int
}

func (s *sorterZCount) Len() int      { return len(s.order) }
func (s *sorterZCount) Swap(i, j int) { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *sorterZCount) Less(i, j int) bool {
	vi := int64(s.get(s.order[i]))
	vj := int64(s.get(s.order[j]))
	return bit.Count(bit.ZEncode(vi)) < bit.Count(bit.ZEncode(vj))
}

func SortByZCount(items []int, get func(i int) int) {
	sort.Sort(&sorterZCount{
		order: items,
		get:   get,
	})
}
