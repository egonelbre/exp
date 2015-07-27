package sse

func AddU32(dst, src []uint32)
func SubU32(dst, src []uint32)
func MulU32(dst, src []uint32)

func MinU32Len(dst, src []uint32) int

func addU32(dst, src []uint32) {
	n := len(dst)
	if n > len(src) {
		n = len(src)
	}

	for i := range dst[:n] {
		dst[i] += src[i]
	}
}

func subU32(dst, src []uint32) {
	n := len(dst)
	if n > len(src) {
		n = len(src)
	}

	for i := range dst[:n] {
		dst[i] -= src[i]
	}
}

func mulU32(dst, src []uint32) {
	n := len(dst)
	if n > len(src) {
		n = len(src)
	}

	for i := range dst[:n] {
		dst[i] *= src[i]
	}
}
