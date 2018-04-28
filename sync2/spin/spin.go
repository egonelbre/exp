package spin

import "runtime"

type T = T256

type T256 struct{ count int }

func (s *T256) Spin() bool {
	s.count++
	if s.count > 256 {
		runtime.Gosched()
		s.count = 0
	}
	return true
}
