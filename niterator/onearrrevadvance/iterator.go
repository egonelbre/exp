package onearrrevadvance

import (
	"io"

	"github.com/egonelbre/exp/niterator/shape"
)

type Index struct {
	Track   uint32
	Advance uint32
	Shape   uint32
}

type Iterator struct {
	*shape.AP

	Track     []Index
	NextIndex uint32
	Done      bool
}

func NewIterator(ap *shape.AP) *Iterator {
	track := make([]Index, len(ap.Shape))

	last := len(ap.Shape) - 1
	stride := ap.Stride[:last+1]
	shape := ap.Shape[:last+1]

	track[0].Track = uint32(shape[last])
	track[0].Advance = uint32(stride[last])
	track[0].Shape = uint32(shape[last])

	for i := 1; i < last+1; i++ {
		track[i].Track = uint32(shape[last-i])
		track[i].Shape = uint32(shape[last-i])
		track[i].Advance = uint32(stride[last-i] - stride[last-i+1]*shape[last-i+1])
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

	next := it.NextIndex
	result := next
	track := it.Track
	for i := range track {
		x := &track[i]
		x.Track--
		next += x.Advance
		if x.Track > 0 {
			it.NextIndex = next
			return int(result), nil
		}
		x.Track = x.Shape
	}

	it.Done = true
	it.NextIndex = next
	return int(result), nil
}
