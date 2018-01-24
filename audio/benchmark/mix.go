package benchmark_audio

import "unsafe"

func Mix_Dynamic(dst, a, b []float32) {
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	if n > len(dst) {
		n = len(dst)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] + b[i]
	}
}

func Mix_Baseline(dst, a, b []float32) {
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	if n > len(dst) {
		n = len(dst)
	}

	pdst := (uintptr)(unsafe.Pointer(&dst[0]))
	pend := (uintptr)(unsafe.Pointer(&dst[len(dst)-1])) + 4
	pa := (uintptr)(unsafe.Pointer(&a[0]))
	pb := (uintptr)(unsafe.Pointer(&b[0]))
	for pdst < pend {
		*(*float32)(unsafe.Pointer(pdst)) =
			*(*float32)(unsafe.Pointer(pa)) +
				*(*float32)(unsafe.Pointer(pb))
		pdst += 4
		pa += 4
		pb += 4
	}
}

func MixInto_Dynamic(a, b []float32) {
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		a[i] += b[i]
	}
}

func MixInto_Baseline(a, b []float32) {
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	pa := (uintptr)(unsafe.Pointer(&a[0]))
	pend := (uintptr)(unsafe.Pointer(&a[len(a)-1])) + 4
	pb := (uintptr)(unsafe.Pointer(&b[0]))
	for pa < pend {
		*(*float32)(unsafe.Pointer(pa)) += *(*float32)(unsafe.Pointer(pb))
		pa += 4
		pb += 4
	}
}
