package shape

type AP struct {
	Shape  []int
	Stride []int
}

func (ap *AP) TotalSize() int {
	total := 1
	for _, size := range ap.Shape {
		total *= size
	}
	return total
}

func New(shape ...int) *AP {
	ap := &AP{}
	ap.Shape = shape
	ap.Stride = make([]int, len(shape))

	acc := 1
	for i := len(shape) - 1; i >= 0; i-- {
		ap.Stride[i] = acc
		d := shape[i]
		if d < 0 {
			panic("negative dimension size does not make sense")
		}
		acc *= d
	}
	return ap
}
