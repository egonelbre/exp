package onearr

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

	Track     []Index
	NextIndex uint32
	Done      bool
}

func NewIterator(ap *shape.AP) *Iterator {
	track := make([]Index, len(ap.Shape))
	for i := range ap.Shape {
		track[i].Shape = uint32(ap.Shape[i])
		track[i].Stride = uint32(ap.Stride[i])
	}

	return &Iterator{
		AP:    ap,
		Track: track,
	}
}

func (it *Iterator) IsDone() bool {
	return it.Done
}

func (it *Iterator) Next() (int, error) {
	if it.Done {
		return 0, io.EOF
	}

	last := len(it.Track) - 1
	next := it.NextIndex
	result := next

	// the following 3 lines causes the compiler to perform bounds check here,
	// instead of being done in the loop
	track := it.Track[:last+1]
	for i := last; i >= 0; i-- {
		x := &track[i]
		x.Track++
		if x.Track == x.Shape {
			if i == 0 {
				it.Done = true
			}
			x.Track = 0
			next -= (x.Shape - 1) * x.Stride
			continue
		}
		next += x.Stride
		break
	}
	it.NextIndex = next
	return int(result), nil
}
