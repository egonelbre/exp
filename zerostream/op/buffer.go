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
	Data []byte
}

type Iterator struct {
	Stream
	Head int
}

func (it *Iterator) EOF() bool { return it.Type() == EOF }

func (it *Iterator) Type() Type { return Type(it.Data[it.Head]) }

func (it *Iterator) Start() *Start {
	it.Data[it.Head] = byte(TypeStart)
	r := (*Start)(unsafe.Pointer(&it.Data[it.Head+1]))
	it.Head += 1 + int(unsafe.Sizeof(Start{}))
	return r
}

func (it *Iterator) Close() *Close {
	it.Data[it.Head] = byte(TypeClose)
	r := (*Close)(unsafe.Pointer(&it.Data[it.Head+1]))
	it.Head += 1 + int(unsafe.Sizeof(Close{}))
	return r
}

func (it *Iterator) Translate() *Translate {
	it.Data[it.Head] = byte(TypeTranslate)
	r := (*Translate)(unsafe.Pointer(&it.Data[it.Head+1]))
	it.Head += 1 + int(unsafe.Sizeof(Translate{}))
	return r
}

func (it *Iterator) MoveTo() *MoveTo {
	it.Data[it.Head] = byte(TypeMoveTo)
	r := (*MoveTo)(unsafe.Pointer(&it.Data[it.Head+1]))
	it.Head += 1 + int(unsafe.Sizeof(MoveTo{}))
	return r
}

func (it *Iterator) LineTo() *LineTo {
	it.Data[it.Head] = byte(TypeLineTo)
	r := (*LineTo)(unsafe.Pointer(&it.Data[it.Head+1]))
	it.Head += 1 + int(unsafe.Sizeof(LineTo{}))
	return r
}
