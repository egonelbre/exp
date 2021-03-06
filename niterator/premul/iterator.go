package premul

import (
	"io"

	"github.com/egonelbre/exp/niterator/shape"
)

type Iterator struct {
	*shape.AP

	Track     []int
	Premul    []int
	NextIndex int
	Done      bool
}

func NewIterator(ap *shape.AP) *Iterator {
	premul := make([]int, len(ap.Shape))
	for i := range ap.Shape {
		premul[i] = (ap.Shape[i] - 1) * ap.Stride[i]
	}

	return &Iterator{
		AP:     ap,
		Premul: premul,
		Track:  make([]int, len(ap.Shape)),
	}
}

func (it *Iterator) IsDone() bool {
	return it.Done
}

func (it *Iterator) Next() (int, error) {
	if it.Done {
		return 0, io.EOF
	}

	last := len(it.Shape) - 1
	next := it.NextIndex
	result := next

	// the following 3 lines causes the compiler to perform bounds check here,
	// instead of being done in the loop
	coord := it.Shape[:last+1]
	track := it.Track[:last+1]
	stride := it.Stride[:last+1]
	premul := it.Premul[:last+1]
	for i := last; i >= 0; i-- {
		track[i]++
		shapeI := coord[i]
		strideI := stride[i]

		if track[i] == shapeI {
			if i == 0 {
				it.Done = true
			}
			track[i] = 0
			next -= premul[i]
			continue
		}
		next += strideI
		break
	}
	it.NextIndex = next
	return result, nil
}
