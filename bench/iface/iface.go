package iface

type Interface interface {
	Get() int
	CopyGet() int
	Set(x int)
}

//go:noinline
func NewStruct(i int) Interface {
	if i == 0 {
		return &Struct{}
	} else {
		return &Struct2{}
	}
}

type Struct struct {
	x int
}

func (s Struct) CopyGet() int {
	return s.x
}

func (s *Struct) Get() int {
	return s.x
}

func (s *Struct) Set(x int) {
	s.x = x
}

type Struct2 struct {
	x int
}

func (s Struct2) CopyGet() int {
	return s.x
}

func (s *Struct2) Get() int {
	return s.x
}

func (s *Struct2) Set(x int) {
	s.x = x
}
