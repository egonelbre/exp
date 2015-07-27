#include "textflag.h"

// func AddU32_SSE(dst, src []uint32)
TEXT ·AddU32_SSE(SB),NOSPLIT,$0-48
	MOVQ     dst+8(FP),  CX // CX = len(dst)
	MOVQ     src+32(FP), BX // BX = len(src)
	CMPQ     CX, BX         //
	CMOVQGT  BX, CX         // CX = min(CX, BX)
	MOVQ     dst+0(FP),  AX // AX = &dst[0]
	MOVQ     src+24(FP), BX // BX = &src[0]

	vector:
		SUBQ     $16, CX
		JL trailing
		// load src
		MOVOU  +00(BX), X0
		MOVOU  +16(BX), X1
		MOVOU  +32(BX), X2
		MOVOU  +48(BX), X3
		// load dst
		MOVOU  +00(AX), X4
		MOVOU  +16(AX), X5
		MOVOU  +32(AX), X6
		MOVOU  +48(AX), X7
		// dst + src
		PADDL X0, X4
		PADDL X1, X5
		PADDL X2, X6
		PADDL X3, X7
		// store to dst
		MOVOU X4, +00(AX)
		MOVOU X5, +16(AX)
		MOVOU X6, +32(AX)
		MOVOU X7, +48(AX)
		// increment
		ADDQ $64, BX
		ADDQ $64, AX
		JMP vector

	trailing:
		ADDQ $17, CX
	elem:
		DECQ     CX
		JZ done
		MOVL (BX), DX
		ADDL DX,  (AX)
		ADDQ $4, BX
		ADDQ $4, AX
		JMP elem
	done:
		RET


// func SubU32_SSE(dst, src []uint32)
TEXT ·SubU32_SSE(SB),NOSPLIT,$0
	MOVQ     dst+8(FP),  CX // CX = len(dst)
	MOVQ     src+32(FP), BX // BX = len(src)
	CMPQ     CX, BX         //
	CMOVQGT  BX, CX         // CX = min(CX, BX)
	MOVQ     dst+0(FP),  AX // AX = &dst[0]
	MOVQ     src+24(FP), BX // BX = &src[0]

	vector:
		SUBQ     $16, CX
		JL trailing
		// load src
		MOVOU  +00(BX), X0
		MOVOU  +16(BX), X1
		MOVOU  +32(BX), X2
		MOVOU  +48(BX), X3
		// load dst
		MOVOU  +00(AX), X4
		MOVOU  +16(AX), X5
		MOVOU  +32(AX), X6
		MOVOU  +48(AX), X7
		// dst - src
		PSUBL X0, X4
		PSUBL X1, X5
		PSUBL X2, X6
		PSUBL X3, X7
		// store to dst
		MOVOU X4, +00(AX)
		MOVOU X5, +16(AX)
		MOVOU X6, +32(AX)
		MOVOU X7, +48(AX)
		// increment
		ADDQ $64, BX
		ADDQ $64, AX
		JMP vector

	trailing:
		ADDQ $17, CX
	elem:
		DECQ     CX
		JZ done
		MOVL (BX), DX
		// sub
		SUBL DX, (AX)
		// increment
		ADDQ $4, BX
		ADDQ $4, AX
		JMP elem
	done:
		RET

// func MulU32_SSE(dst, src []uint32)
TEXT ·MulU32_SSE(SB),NOSPLIT,$0
	MOVQ     dst+8(FP),  CX // CX = len(dst)
	MOVQ     src+32(FP), BX // BX = len(src)
	CMPQ     CX, BX         //
	CMOVQGT  BX, CX         // CX = min(CX, BX)
	MOVQ     dst+0(FP),  AX // AX = &dst[0]
	MOVQ     src+24(FP), BX // BX = &src[0]

	vector:
		SUBQ     $16, CX
		JL trailing
		// load src
		MOVOU  +00(BX), X0
		MOVOU  +16(BX), X1
		MOVOU  +32(BX), X2
		MOVOU  +48(BX), X3
		// load dst
		MOVOU  +00(AX), X4
		MOVOU  +16(AX), X5
		MOVOU  +32(AX), X6
		MOVOU  +48(AX), X7
		// dst - src
		PSUBL X0, X4
		PSUBL X1, X5
		PSUBL X2, X6
		PSUBL X3, X7
		// store to dst
		MOVOU X4, +00(AX)
		MOVOU X5, +16(AX)
		MOVOU X6, +32(AX)
		MOVOU X7, +48(AX)
		// increment
		ADDQ $64, BX
		ADDQ $64, AX
		JMP vector

	trailing:
		ADDQ $17, CX
	elem:
		DECQ     CX
		JZ done
		MOVL (BX), DX
		// sub
		IMULL (AX), DX
		MOVL  DX, (AX)
		// increment
		ADDQ $4, BX
		ADDQ $4, AX
		JMP elem
	done:
		RET
