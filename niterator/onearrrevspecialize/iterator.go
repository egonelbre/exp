package onearrrevspecialize

import (
	"io"

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
	it := &Iterator{}
	it.AP = ap

	last := len(ap.Shape) - 1
	track := &it.Track
	for i := range ap.Shape {
		(*track)[last-i].Track = uint32(ap.Shape[i])
		(*track)[last-i].Shape = uint32(ap.Shape[i])
		(*track)[last-i].Stride = uint32(ap.Stride[i])
	}

	return it
}

func (it *Iterator) IsDone() bool {
	return it.Done
}

func (it *Iterator) Next() (int, error) {
	if it.Done {
		return 0, io.EOF
	}

	next := it.NextIndex
	result := next
	for i := range it.Track {
		x := &it.Track[i]
		x.Track--
		if x.Track > 0 {
			next += x.Stride
			it.NextIndex = next
			return int(result), nil
		}
		x.Track = x.Shape
		next -= (x.Shape - 1) * x.Stride
	}

	it.Done = true
	it.NextIndex = next
	return int(result), nil
}
