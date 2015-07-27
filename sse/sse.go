package sse

func AddU32_ASM(dst, src []uint32)
func SubU32_ASM(dst, src []uint32)
func MulU32_ASM(dst, src []uint32)

func AddU32_Slow(dst, src []uint32) {
	n := len(dst)
	if n > len(src) {
		n = len(src)
	}

	for i := range dst[:n] {
		dst[i] += src[i]
	}
}

func SubU32_Slow(dst, src []uint32) {
	n := len(dst)
	if n > len(src) {
		n = len(src)
	}

	for i := range dst[:n] {
		dst[i] -= src[i]
	}
}

func MulU32_Slow(dst, src []uint32) {
	n := len(dst)
	if n > len(src) {
		n = len(src)
	}

	for i := range dst[:n] {
		dst[i] *= src[i]
	}
}
