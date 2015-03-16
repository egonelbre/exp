package main

type IndexValue struct {
	Value byte
	Index int
}

func (v IndexValue) Get(deltas Deltas) int32 {
	switch v.Value {
	case 0:
		return deltas[v.Index].A
	case 1:
		return deltas[v.Index].B
	case 2:
		return deltas[v.Index].C
	case 3:
		return deltas[v.Index].X
	case 4:
		return deltas[v.Index].Y
	case 5:
		return deltas[v.Index].Z
	}
	return 0
}

func (v IndexValue) Set(deltas Deltas, x int32) {
	switch v.Value {
	case 0:
		deltas[v.Index].A = x
	case 1:
		deltas[v.Index].B = x
	case 2:
		deltas[v.Index].C = x
	case 3:
		deltas[v.Index].X = x
	case 4:
		deltas[v.Index].Y = x
	case 5:
		deltas[v.Index].Z = x
	}
}

type byValue struct {
	order  []IndexValue
	deltas Deltas
}

func (s *byValue) Len() int           { return len(s.order) }
func (s *byValue) Swap(i, j int)      { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *byValue) Less(i, j int) bool { return s.order[i].Get(s.deltas) < s.order[j].Get(s.deltas) }

type byDelta struct {
	order    []IndexValue
	baseline Deltas
	current  Deltas
}

func (s *byDelta) Len() int      { return len(s.order) }
func (s *byDelta) Swap(i, j int) { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *byDelta) Less(i, j int) bool {
	di := s.order[i].Get(s.baseline) - s.order[i].Get(s.current)
	dj := s.order[j].Get(s.baseline) - s.order[j].Get(s.current)
	return di < dj
}

func sequence(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}

func NewOrdering(deltas Deltas) *Ordering {
	n := len(deltas)
	order := &Ordering{
		Largest:     sequence(n),
		Interacting: sequence(n),

		ABC: make([]IndexValue, n*3),
		XYZ: make([]IndexValue, n*3),
	}

	for i := 0; i < n*3; i += 1 {
		order.ABC[i].Index = i % n
		order.ABC[i].Value = byte(i / n)

		order.XYZ[i].Index = i % n
		order.XYZ[i].Value = byte(i/n + 3)
	}

	return order
}

type sorter struct {
	order  []int
	deltas Deltas
	get    Getter
}

func (s *sorter) Len() int      { return len(s.order) }
func (s *sorter) Swap(i, j int) { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *sorter) Less(i, j int) bool {
	return s.get(&s.deltas[s.order[i]]) < s.get(&s.deltas[s.order[j]])
}
