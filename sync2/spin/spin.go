package spin

import "runtime"

// Default spintter
type T = T256

// Default limited spinner
type L = L256_256

// T256 is a spinner with Gosched every 256 iterations
type T256 struct{ count int }

func (s *T256) Spin() bool {
	s.count++
	if s.count > 256 {
		runtime.Gosched()
		s.count = 0
	}
	return true
}

// L256_256 is a limited spinner with Gosched every 256 iteration and maximum 1024 Gosched-s
type L256_256 struct {
	count int
	total int
}

func (s *L256_256) Spin() bool {
	s.count++
	if s.count > 256 {
		runtime.Gosched()
		s.count = 0
		s.total++
	}
	if s.total > 256 {
		return false
	}
	return true
}

// Limited spinner with Gosched every 128 iteration and maximum 1024 Gosched-s
type L128_1024 struct {
	count int
	total int
}

func (s *L128_1024) Spin() bool {
	s.count++
	if s.count > 128 {
		runtime.Gosched()
		s.count = 0
		s.total++
	}
	if s.total > 1024 {
		return false
	}
	return true
}
