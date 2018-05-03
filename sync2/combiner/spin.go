package combiner

import "runtime"

func spin(v *int) {
	*v++
	if *v >= 128 {
		runtime.Gosched()
	}
}
