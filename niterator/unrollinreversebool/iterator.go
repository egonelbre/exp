package unrollinreversebool

import (
	"github.com/egonelbre/exp/niterator/shape"
)

type Index struct {
	Track  uint32
	Shape  uint32
	Stride uint32
}

type Iterator struct {
	*shape.AP

	Track     [4]Index
	NextIndex uint32
	Done      bool
}

func NewIterator(ap *shape.AP) *Iterator {
	track := [4]Index{}
	for i := range ap.Shape {
		track[3-i].Track = uint32(ap.Shape[i])
		track[3-i].Shape = uint32(ap.Shape[i])
		track[3-i].Stride = uint32(ap.Stride[i])
	}

	return &Iterator{
		AP:    ap,
		Track: track,
	}
}

func (it *Iterator) IsDone() bool {
	return it.Done
}

func (it *Iterator) Next() (int, bool) {
	if it.Done {
		return 0, it.Done
	}

	next := it.NextIndex
	result := int(next)

	x := &it.Track[0]
	x.Track--
	if x.Track > 0 {
		next += x.Stride
		it.NextIndex = next
		return result, false
	}
	x.Track = x.Shape
	next -= (x.Shape - 1) * x.Stride

	x = &it.Track[1]
	x.Track--
	if x.Track > 0 {
		next += x.Stride
		it.NextIndex = next
		return result, false
	}
	x.Track = x.Shape
	next -= (x.Shape - 1) * x.Stride

	x = &it.Track[2]
	x.Track--
	if x.Track > 0 {
		next += x.Stride
		it.NextIndex = next
		return result, false
	}
	x.Track = x.Shape
	next -= (x.Shape - 1) * x.Stride

	x = &it.Track[3]
	x.Track--
	if x.Track > 0 {
		next += x.Stride
		it.NextIndex = next
		return result, false
	}
	it.Done = true
	it.NextIndex = next

	return result, false
}
