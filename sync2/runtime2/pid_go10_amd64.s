#include "textflag.h"
#include "go_asm.h"
#include "go_tls.go110.h"

TEXT ·asm_pid(SB),NOSPLIT,$0-4
	get_tls(CX)
	MOVQ	g(CX), AX
	MOVQ    G_M_OFFSET(AX), AX
	MOVQ    M_P_OFFSET(AX), AX
	MOVL    P_ID_OFFSET(AX), AX
	MOVL	AX, ret+0(FP)
	RET
