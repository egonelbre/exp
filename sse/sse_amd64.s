#include "textflag.h"

#define Load4x(FROM, A, B, C, D) \
	MOVOU  +00(FROM), A; \
	MOVOU  +16(FROM), B; \
	MOVOU  +32(FROM), C; \
	MOVOU  +48(FROM), D;

#define Store4x(INTO, A, B, C, D)  \
	MOVOU  A, +00(INTO); \
	MOVOU  B, +16(INTO); \
	MOVOU  C, +32(INTO); \
	MOVOU  D, +48(INTO);

#define Apply4x(OP) \
	OP X0, X4;      \
	OP X1, X5;      \
	OP X2, X6;      \
	OP X3, X7;


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

		Load4x(BX, X0, X1, X2, X3)
		Load4x(AX, X4, X5, X6, X7)
		Apply4x(PADDL)
		Store4x(AX, X4, X5, X6, X7)

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

		Load4x(BX, X0, X1, X2, X3)
		Load4x(AX, X4, X5, X6, X7)
		Apply4x(PSUBL)
		Store4x(AX, X4, X5, X6, X7)

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

		Load4x(BX, X0, X1, X2, X3)
		Load4x(AX, X4, X5, X6, X7)
		Apply4x(PMULLW)
		Store4x(AX, X4, X5, X6, X7)

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
