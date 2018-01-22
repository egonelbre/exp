package instruct

import (
	"io"

	"github.com/egonelbre/exp/niterator/shape"
)

type Iterator struct {
	*shape.AP

	UShape  []uint32
	UStride []uint32
	UTrack  []uint32

	NextIndex uint32
	Done      bool
}

func NewIterator(ap *shape.AP) *Iterator {
	xshape := make([]uint32, len(ap.Shape))
	xstride := make([]uint32, len(ap.Stride))

	for i, v := range ap.Shape {
		xshape[i] = uint32(v)
	}
	for i, v := range ap.Stride {
		xstride[i] = uint32(v)
	}

	return &Iterator{
		AP:      ap,
		UShape:  xshape,
		UStride: xstride,
		UTrack:  make([]uint32, len(ap.Shape)),
	}
}

func (it *Iterator) IsDone() bool {
	return it.Done
}

func (it *Iterator) Next() (int, error) {
	if it.Done {
		return 0, io.EOF
	}

	last := len(it.UShape) - 1
	next := it.NextIndex
	result := next

	// the following 3 lines causes the compiler to perform bounds check here,
	// instead of being done in the loop
	coord := it.UShape[:last+1]
	track := it.UTrack[:last+1]
	stride := it.UStride[:last+1]
	for i := last; i >= 0; i-- {
		track[i]++
		shapeI := coord[i]
		strideI := stride[i]

		if track[i] == shapeI {
			if i == 0 {
				it.Done = true
			}
			track[i] = 0
			next -= (shapeI - 1) * strideI
			continue
		}
		next += strideI
		break
	}
	it.NextIndex = next
	return int(result), nil
}
