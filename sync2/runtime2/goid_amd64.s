// +build go1.10

#include "textflag.h"
#include "go_asm.h"
#include "go_tls.go110.h"

// func GOID() int64
TEXT ·GOID(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ	g(CX), DX
	MOVQ    G_ID_OFFSET(DX), AX
	MOVQ	AX, ret+0(FP)
	RET

// func g() uintptr
TEXT ·g(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ	g(CX), BX
	MOVQ	BX, ret+0(FP)
	RET
