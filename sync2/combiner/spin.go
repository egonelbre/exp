package combiner

import "runtime"

type pad7 [7]uint64

func spin(v *int) {
	*v++
	if *v >= 128 {
		runtime.Gosched()
	}
}
