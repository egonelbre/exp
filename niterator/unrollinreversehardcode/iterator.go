package unrollinreversehardcode

import (
	"io"

	"github.com/egonelbre/exp/niterator/shape"
)

type Index struct {
	Track uint32
	//Shape  uint32
	//Stride uint32
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
		//track[3-i].Shape = uint32(ap.Shape[i])
		//track[3-i].Stride = uint32(ap.Stride[i])
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
	result := int(next)

	// {Track:100 Shape:100 Stride:1}
	x := &it.Track[0]
	x.Track--
	next += 1
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = 100
	next -= 100 * 1

	// {Track:100 Shape:100 Stride:100}
	x = &it.Track[1]
	x.Track--
	next += 100
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = 100
	next -= 100 * 100

	// {Track:100 Shape:100 Stride:10000}
	x = &it.Track[2]
	x.Track--
	next += 10000
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = 100
	next -= 100 * 10000

	// {Track:100 Shape:100 Stride:1000000}
	x = &it.Track[3]
	x.Track--
	next += 1000000
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	it.Done = true
	it.NextIndex = next

	return result, nil
}
