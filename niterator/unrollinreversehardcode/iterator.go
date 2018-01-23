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

// 30, 30, 1; 40 40 30; 20 20 1200; 24 24 24000

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

	// {Track:30 Shape:30 Stride:1}
	x := &it.Track[0]
	x.Track--
	next += 1
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = 30
	next -= 30 * 1

	// {Track:40 Shape:40 Stride:30}
	x = &it.Track[1]
	x.Track--
	next += 30
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = 40
	next -= 40 * 30

	// {Track:20 Shape:20 Stride:1200}
	x = &it.Track[2]
	x.Track--
	next += 1200
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = 20
	next -= 20 * 1200

	// {Track:24 Shape:24 Stride:24000}
	x = &it.Track[3]
	x.Track--
	next += 24000
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	it.Done = true
	it.NextIndex = next

	return result, nil
}
