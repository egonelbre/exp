package queue

import "runtime"

func spin(v *int) {
	*v++
	if *v&0xFF == 0 {
		runtime.Gosched()
	}
}
