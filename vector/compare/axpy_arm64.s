#include "textflag.h"

// func ArmAxpyUnsafe(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·ArmAxpyUnsafe(SB), NOSPLIT, $0-48
	FMOVS alpha+0(FP), F0

	MOVD  xs+8(FP),    R0
	MOVD  incx+16(FP), R1

	MOVD  ys+24(FP),   R2
	MOVD  incy+32(FP), R3

	MOVD  n+40(FP), R4

	MOVD ZR, R5 // i
	MOVD ZR, R6 // xi
	MOVD ZR, R7 // yi

	JMP   check_limit

loop:
	FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	FMADDS F0, F2, F1, F2  // f1 = alpha * xs[xi] + ys[yi]
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	ADD  $1, R5, R5 // i++
	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy

check_limit:
	CMP  R5, R4
	BHI  loop
	RET

// func ArmAxpyUnsafeX(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·ArmAxpyUnsafeX(SB), NOSPLIT, $0-48
	FMOVS alpha+0(FP), F0

	MOVD  xs+8(FP),    R0
	MOVD  incx+16(FP), R1

	MOVD  ys+24(FP),   R2
	MOVD  incy+32(FP), R3

	MOVD  n+40(FP), R4

	MOVD ZR, R6 // xi
	MOVD ZR, R7 // yi

	JMP   check_limit

loop:
	FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	// FMADDS A, B, C, D ==> D := A*C + B
	FMADDS F0, F2, F1, F2  // f1 = alpha * xs[xi] + ys[yi]
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	SUB  $1, R4, R4 // n--
	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy

check_limit:
	CBNZ R4, loop
	RET

// func ArmAxpyPointer(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·ArmAxpyPointer(SB), NOSPLIT, $0-48
	FMOVS alpha+0(FP), F0

	MOVD  xs+8(FP),    R0
	MOVD  incx+16(FP), R1

	MOVD  ys+24(FP),   R2
	MOVD  incy+32(FP), R3

	MOVD  n+40(FP), R4
	LSL   $2, R4, R4
	MADD  R4, R0, R1, R4  // MADD A, B, C, D ==> D := A*C + B

	JMP   check_limit

loop:
	FMOVS  (R0), F1        // f1 = *xs
	FMOVS  (R2), F2        // f2 = *ys
	FMADDS F0, F2, F1, F2  // f1 = alpha * xs[xi] + ys[yi]
	FMOVS  F2, (R2)        // ys[yi] = f1

	ADD  R1<<2, R0, R0 // xs++
	ADD  R3<<2, R2, R2 // ys++

check_limit:
	CMP  R0, R4
	BHI  loop
	RET

// func ArmAxpyPointerLoop(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·ArmAxpyPointerLoop(SB), NOSPLIT, $0-48
	FMOVS alpha+0(FP), F0

	MOVD  xs+8(FP),    R0
	MOVD  incx+16(FP), R1

	MOVD  ys+24(FP),   R2
	MOVD  incy+32(FP), R3

	MOVD  n+40(FP), R4
	MOVD  ZR, R5

	JMP   check_limit

loop:
	FMOVS  (R0), F1        // f1 = *xs
	FMOVS  (R2), F2        // f2 = *ys
	FMADDS F0, F2, F1, F2  // f1 = alpha * xs[xi] + ys[yi]
	FMOVS  F2, (R2)        // ys[yi] = f1

	ADD  $1, R5, R5
	ADD  R1<<2, R0, R0 // xs++
	ADD  R3<<2, R2, R2 // ys++

check_limit:
	CMP  R5, R4
	BHI  loop
	RET

// func ArmAxpyPointerLoopX(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·ArmAxpyPointerLoopX(SB), NOSPLIT, $0-48
	FMOVS alpha+0(FP), F0

	MOVD  xs+8(FP),    R0
	MOVD  incx+16(FP), R1

	MOVD  ys+24(FP),   R2
	MOVD  incy+32(FP), R3

	MOVD  n+40(FP), R4

	JMP   check_limit

loop:
	FMOVS  (R0), F1        // f1 = *xs
	FMOVS  (R2), F2        // f2 = *ys
	FMADDS F0, F2, F1, F2  // f1 = alpha * xs[xi] + ys[yi]
	FMOVS  F2, (R2)        // ys[yi] = f1

	SUB  $1, R4, R4
	ADD  R1<<2, R0, R0 // xs++
	ADD  R3<<2, R2, R2 // ys++

check_limit:
	CBNZ  R4, loop
	RET

// func ArmAxpyUnsafeXR4(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
TEXT ·ArmAxpyUnsafeXR4(SB), NOSPLIT, $0-48
	FMOVS alpha+0(FP), F0

	MOVD  xs+8(FP),    R0
	MOVD  incx+16(FP), R1

	MOVD  ys+24(FP),   R2
	MOVD  incy+32(FP), R3

	MOVD  n+40(FP), R4 // n

	MOVD ZR, R6 // xi
	MOVD ZR, R7 // yi

	JMP  check_limit_unroll

loop_unroll:
    FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	FMADDS F0, F2, F1, F2  // f1 = alpha * f1 + f2
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy
    FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	FMADDS F0, F2, F1, F2  // f1 = alpha * f1 + f2
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy
    FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	FMADDS F0, F2, F1, F2  // f1 = alpha * f1 + f2
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy
    FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	FMADDS F0, F2, F1, F2  // f1 = alpha * f1 + f2
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy
	SUB  $4, R4, R4    // n -= 4

check_limit_unroll:
	CMP  $0x04, R4
	BHS  loop_unroll
	JMP  check_limit

loop:
	FMOVS  (R0)(R6<<2), F1 // f1 = xs[xi]
	FMOVS  (R2)(R7<<2), F2 // f2 = ys[yi]
	// FMADDS A, B, C, D ==> D := A*C + B
	FMADDS F0, F2, F1, F2  // f1 = alpha * xs[xi] + ys[yi]
	FMOVS  F2, (R2)(R7<<2) // ys[yi] = f1

	SUB  $1, R4, R4 // n--
	ADD  R6, R1, R6 // xi += incx
	ADD  R7, R3, R7 // yi += incy

check_limit:
	CBNZ R4, loop
	RET
