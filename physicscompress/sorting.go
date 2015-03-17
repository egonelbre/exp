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

type byGetter struct {
	order  []int
	deltas Deltas
	get    Getter
}

func (s *byGetter) Len() int      { return len(s.order) }
func (s *byGetter) Swap(i, j int) { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *byGetter) Less(i, j int) bool {
	return s.get(&s.deltas[s.order[i]]) < s.get(&s.deltas[s.order[j]])
}
