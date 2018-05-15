package iface

type Interface interface {
	GetX() uint
	GetXIndirect() uint
	SetX(x uint)
}
type Struct struct {
	x uint
}

func (s Struct) GetXIndirect() uint {
	return s.x
}

func (s *Struct) GetX() uint {
	return s.x
}

func (s *Struct) SetX(x uint) {
	s.x = x
}
