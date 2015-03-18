package physics

import (
	"flag"
	"sort"
)

type Values interface {
	Index(i int) int

	// number of values in this projection
	Len() int
	Get(i int) int32
	Set(i int, v int32)
}

var flagSkipOrdering = flag.Bool("order", true, "whether to adjust item serialization ordering")

func ordering(vals Values) []int {
	order := make([]int, vals.Len())
	for i := range order {
		order[i] = i
	}
	if *flagSkipOrdering {
		sort.Sort(&byValues{order, vals})
	}
	return order
}

type byValues struct {
	Order  []int
	Values Values
}

func (s *byValues) Len() int           { return len(s.Order) }
func (s *byValues) Swap(i, j int)      { s.Order[i], s.Order[j] = s.Order[j], s.Order[i] }
func (s *byValues) Less(i, j int) bool { return s.Values.Get(i) < s.Values.Get(j) }

type Delta struct{ Base, Next Values }

// returns the index of the cube
func (delta Delta) Index(i int) int { return delta.Base.Index(i) }

func (delta Delta) Len() int           { return delta.Base.Len() }
func (delta Delta) Get(i int) int32    { return delta.Next.Get(i) - delta.Base.Get(i) }
func (delta Delta) Set(i int, v int32) { delta.Next.Set(i, delta.Base.Get(i)+v) }
