package combiner

import "runtime"

type spinner struct{ count int }

func (s *spinner) spin() bool {
	s.count++
	if s.count > 256 {
		runtime.Gosched()
		s.count = 0
	}
	return true
}
