package main

import (
	"fmt"

	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/ir"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {
	AxpyPointer(11)
	Generate()
}

func AxpyPointer(align int) {
	TEXT(fmt.Sprintf("AxpyPointer_%v", align), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	end := n
	SHLQ(U8(0x2), end)
	IMULQ(incx, end)
	ADDQ(xs.Base, end)
	JMP(LabelRef("check_limit"))

	Align(align)
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs, tmp)
		MULSS(alpha, tmp)
		ADDSS(ys, tmp)
		MOVSS(tmp, ys)

		xs = xs.Idx(incx, 4)
		ys = ys.Idx(incy, 4)
		Label("check_limit")

		CMPQ(end, xs)
		JHI(LabelRef("loop"))
	}

	RET()
}

func Align(n int) {
	if n < 8 {
		panic(fmt.Sprint("alignment %v not supported", n))
	}

	nearestPowerOf2 := 8
	for n >= nearestPowerOf2*2 {
		nearestPowerOf2 *= 2
	}

	Instruction(&ir.Instruction{
		Opcode:   "PCALIGN",
		Operands: []Op{Imm(uint64(nearestPowerOf2))},
	})

	n -= nearestPowerOf2
	for i := 0; i < n; i++ {
		NOP()
	}
}
