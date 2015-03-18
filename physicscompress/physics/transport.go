package physics

type Transport interface {
	Prepare(baseline, current *Frame)
	Bools(same []bool)
	ValueBits(same []bool, order []int, vals Values, bits uint64)
	Values(same []bool, order []int, vals Values)
	Close()
}

func (s *State) Transport(t Transport) {
	historic, baseline, current := s.Prev(7), s.Prev(6), s.Current()

	t.Prepare(baseline, current)

	same := make([]bool, len(current.Cubes))
	for i, cube := range current.Cubes {
		same[i] = cube == baseline.Cubes[i]
	}

	t.Bools(same)

	var order []int

	order = ordering(baseline.Largest())
	t.ValueBits(same, order, current.Largest(), 2)

	order = ordering(baseline.Largest())
	t.ValueBits(same, order, current.Interacting(), 1)

	odabc := Delta{historic.ABC(), baseline.ABC()}
	dabc := Delta{baseline.ABC(), current.ABC()}
	t.Values(same, ordering(odabc), dabc)

	oxyz := Delta{historic.XYZ(), baseline.XYZ()}
	xyz := Delta{baseline.XYZ(), current.XYZ()}
	t.Values(same, ordering(oxyz), xyz)

	t.Close()
}
