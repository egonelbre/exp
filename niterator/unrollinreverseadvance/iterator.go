package unrollinreverseadvance

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

	Track     [4]Index
	NextIndex uint32
	Done      bool
}

func NewIterator(ap *shape.AP) *Iterator {
	//track := [4]Index{}

	// this init loop is slow
	//for i := range ap.Shape {
	//	track[3-i].Track = uint32(ap.Shape[i])
	//	track[3-i].Shape = uint32(ap.Shape[i])
	//	if i == len(ap.Shape)-1 {
	//		track[3-i].Advance = uint32(ap.Stride[i])
	//	} else {
	//		track[3-i].Advance = uint32(ap.Stride[i] - ap.Stride[i+1]*ap.Shape[i+1])
	//	}
	//}
	//fmt.Println(track)
	track := [4]Index{Index{30, 1, 30}, Index{40, 0, 40}, Index{20, 0, 20}, Index{24, 0, 24}}

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

	x := &it.Track[0]
	x.Track--
	next += x.Advance
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = x.Shape

	x = &it.Track[1]
	x.Track--
	next += x.Advance
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = x.Shape

	x = &it.Track[2]
	x.Track--
	next += x.Advance
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	x.Track = x.Shape

	x = &it.Track[3]
	x.Track--
	next += x.Advance
	if x.Track > 0 {
		it.NextIndex = next
		return result, nil
	}
	it.Done = true
	it.NextIndex = next

	return result, nil
}
