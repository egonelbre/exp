#include "textflag.h"

DATA  mask51bits<>+0x00(SB)/8, $0x7ffffffffffff
GLOBL mask51bits<>(SB), (NOPTR+RODATA), $8

TEXT ·MaskAndConst(SB),0,$0-0
	MOVQ   $65536,  CX
	MOVQ   $218643, AX

	PCALIGN $32
loop:
	MOVQ   $0x7ffffffffffff, BX
	ANDQ   BX, AX

	DECQ    CX
	JNZ     loop

	RET

TEXT ·MaskShlShr(SB),0,$0-0
	MOVQ   $65536,  CX
	MOVQ   $218643, AX

	PCALIGN $32
loop:
	SHLQ   $13, AX
	SHRQ   $13, AX

	DECQ    CX
	JNZ     loop

	RET

TEXT ·MaskAndAddr(SB),0,$0-0
	MOVQ   $65536,  CX
	MOVQ   $218643, AX

	PCALIGN $32
loop:
	ANDQ   mask51bits<>(SB), AX

	DECQ    CX
	JNZ     loop

	RET
