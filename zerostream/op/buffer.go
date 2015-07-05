package op

import "unsafe"

type Type byte

const (
	EOF = Type(iota)
	TypeStart
	TypeClose
	TypeMoveTo
	TypeLineTo
	TypeStroke
	TypeFill
	TypeLineWidth
	TypeTranslate
	_TypeLength
)

type Start struct{}
type Close struct{}

type Translate struct{ Dx, Dy int32 }

type MoveTo struct{ X, Y int32 }
type LineTo struct{ X, Y int32 }

type LineWidth struct{ Width int32 }
type Stroke struct{}
type Fill struct{}

type Stream struct {
	Base uintptr
	Data []byte
}

func NewStream(max int) Stream {
	data := make([]byte, max, max)
	return Stream{
		Base: uintptr(unsafe.Pointer(&data[0])),
		Data: data,
	}
}

func (s Stream) Iterate() *Iterator {
	return &Iterator{s.Base}
}

type Iterator struct {
	Head uintptr
}

func (it *Iterator) EOF() bool { return it.Type() == EOF }

func (it *Iterator) Type() Type { return Type(*(*byte)(unsafe.Pointer(it.Head))) }

func (it *Iterator) Start() *Start {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeStart)
	r := (*Start)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(Start{})
	return r
}

func (it *Iterator) Close() *Close {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeClose)
	r := (*Close)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(Close{})
	return r
}

func (it *Iterator) Translate() *Translate {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeTranslate)
	r := (*Translate)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(Translate{})
	return r
}

func (it *Iterator) MoveTo() *MoveTo {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeMoveTo)
	r := (*MoveTo)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(MoveTo{})
	return r
}

func (it *Iterator) LineTo() *LineTo {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeLineTo)
	r := (*LineTo)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(LineTo{})
	return r
}

func (it *Iterator) LineWidth() *LineWidth {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeLineWidth)
	r := (*LineWidth)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(LineWidth{})
	return r
}

func (it *Iterator) Stroke() *Stroke {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeStroke)
	r := (*Stroke)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(Stroke{})
	return r
}

func (it *Iterator) Fill() *Fill {
	*(*byte)(unsafe.Pointer(it.Head)) = byte(TypeFill)
	r := (*Fill)(unsafe.Pointer(it.Head + 1))
	it.Head += 1 + unsafe.Sizeof(Fill{})
	return r
}
