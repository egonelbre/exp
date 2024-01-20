// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noasm,!gccgo,!safe

#include "textflag.h"

// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11, R12, R13

// func AsmAxpyPointer_Align11(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyPointer_Align11(SB), NOSPLIT, $0
	MOVSS alpha+0(FP), X0

	MOVQ  xs+8(FP),    AX
	MOVQ  incx+16(FP), BX

	MOVQ  ys+24(FP),   CX
	MOVQ  incy+32(FP), DX

	MOVQ  n+40(FP), DI

	SHLQ  $2, DI // n *= 4
	IMULQ BX, DI // n *= incx
	ADDQ  AX, DI // n += xs

	JMP check_limit
	PCALIGN $8
	NOP
	NOP
	NOP
multiply_loop:
	MOVSS (AX), X1
	MULSS X0, X1
	ADDSS (CX), X1
	MOVSS X1, (CX)

	LEAQ (AX)(BX*4), AX
	LEAQ (CX)(DX*4), CX
check_limit:
	CMPQ DI, AX
	JHI multiply_loop
	RET


// func AsmAxpyPointer_Align16(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyPointer_Align16(SB), NOSPLIT, $0
	MOVSS alpha+0(FP), X0

	MOVQ  xs+8(FP),    AX
	MOVQ  incx+16(FP), BX

	MOVQ  ys+24(FP),   CX
	MOVQ  incy+32(FP), DX

	MOVQ  n+40(FP), DI

	SHLQ  $2, DI // n *= 4
	IMULQ BX, DI // n *= incx
	ADDQ  AX, DI // n += xs

	JMP check_limit
	PCALIGN $16
multiply_loop:
	MOVSS (AX), X1
	MULSS X0, X1
	ADDSS (CX), X1
	MOVSS X1, (CX)

	LEAQ (AX)(BX*4), AX
	LEAQ (CX)(DX*4), CX
check_limit:
	CMPQ DI, AX
	JHI multiply_loop
	RET

// func AsmAxpyPointerLoop_Align11(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyPointerLoop_Align11(SB), NOSPLIT, $0
	MOVSS alpha+0(FP), X0

	MOVQ  xs+8(FP),    AX
	MOVQ  incx+16(FP), BX

	MOVQ  ys+24(FP),   CX
	MOVQ  incy+32(FP), DX

	MOVQ  n+40(FP), DI
	XORQ R12, R12

	JMP check_limit
	PCALIGN $8
	NOP
	NOP
	NOP
multiply_loop:
	MOVSS (AX), X1
	MULSS X0, X1
	ADDSS (CX), X1
	MOVSS X1, (CX)

	INCQ R12

	LEAQ (AX)(BX*4), AX
	LEAQ (CX)(DX*4), CX
check_limit:
	CMPQ DI, R12
	JHI multiply_loop
	RET

// func AsmAxpyPointerLoop_Align16(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyPointerLoop_Align16(SB), NOSPLIT, $0
	MOVSS alpha+0(FP), X0

	MOVQ  xs+8(FP),    AX
	MOVQ  incx+16(FP), BX

	MOVQ  ys+24(FP),   CX
	MOVQ  incy+32(FP), DX

	MOVQ  n+40(FP), DI
	XORQ R12, R12

	JMP check_limit
	PCALIGN $16
multiply_loop:
	MOVSS (AX), X1
	MULSS X0, X1
	ADDSS (CX), X1
	MOVSS X1, (CX)

	INCQ R12

	LEAQ (AX)(BX*4), AX
	LEAQ (CX)(DX*4), CX
check_limit:
	CMPQ DI, R12
	JHI multiply_loop
	RET


// func AsmAxpyUnsafe_Align11(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyUnsafe_Align11(SB), NOSPLIT, $0
	MOVSS alpha+0(FP), X0

	MOVQ  xs+8(FP),    AX
	MOVQ  incx+16(FP), BX

	MOVQ  ys+24(FP),   CX
	MOVQ  incy+32(FP), DX

	MOVQ  n+40(FP), DI
	XORQ R11, R11
	XORQ R12, R12
	XORQ R13, R13

	JMP check_limit
	PCALIGN $8
	NOP
	NOP
	NOP
multiply_loop:
	MOVSS (AX)(R11*4), X1
	MULSS X0, X1
	ADDSS (CX)(R12*4), X1
	MOVSS X1, (CX)

	INCQ R13
	ADDQ BX, R11
	ADDQ DX, R12

check_limit:
	CMPQ DI, R13
	JHI multiply_loop
	RET

// func AsmAxpyUnsafe_Align16(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyUnsafe_Align16(SB), NOSPLIT, $0
	MOVSS alpha+0(FP), X0

	MOVQ  xs+8(FP),    AX
	MOVQ  incx+16(FP), BX

	MOVQ  ys+24(FP),   CX
	MOVQ  incy+32(FP), DX

	MOVQ  n+40(FP), DI
	XORQ R11, R11
	XORQ R12, R12
	XORQ R13, R13

	JMP check_limit
	PCALIGN $16
multiply_loop:
	MOVSS (AX)(R11*4), X1
	MULSS X0, X1
	ADDSS (CX)(R12*4), X1
	MOVSS X1, (CX)

	INCQ R13
	ADDQ BX, R11
	ADDQ DX, R12

check_limit:
	CMPQ DI, R13
	JHI multiply_loop
	RET
