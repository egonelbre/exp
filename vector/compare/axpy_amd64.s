// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noasm,!gccgo,!safe

#include "textflag.h"

// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11

// func AsmAxpyPointer_Align5(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·AsmAxpyPointer_Align5(SB), NOSPLIT, $0
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
