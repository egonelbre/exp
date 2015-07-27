#include "textflag.h"

#define Load4xLow(FROM)   \
	MOVOU  +00(FROM), X0; \
	MOVOU  +16(FROM), X1; \
	MOVOU  +32(FROM), X2; \
	MOVOU  +48(FROM), X3;

#define Store4xLow(INTO)  \
	MOVOU  X0, +00(INTO); \
	MOVOU  X1, +16(INTO); \
	MOVOU  X2, +32(INTO); \
	MOVOU  X3, +48(INTO);

#define Load4xHigh(FROM)  \
	MOVOU  +00(FROM), X4; \
	MOVOU  +16(FROM), X5; \
	MOVOU  +32(FROM), X6; \
	MOVOU  +48(FROM), X7;

#define Store4xHigh(INTO) \
	MOVOU  X4, +00(INTO); \
	MOVOU  X5, +16(INTO); \
	MOVOU  X6, +32(INTO); \
	MOVOU  X7, +48(INTO);

#define Apply4xLowHigh(OP) \
	OP X0, X4;             \
	OP X1, X5;             \
	OP X2, X6;             \
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

		Load4xLow(BX)
		Load4xHigh(AX)
		Apply4xLowHigh(PADDL)
		Store4xHigh(AX)

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

		Load4xLow(BX)
		Load4xHigh(AX)
		Apply4xLowHigh(PSUBL)
		Store4xHigh(AX)

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

		// incorrect code!!!!!!!!
		Load4xLow(BX)
		Load4xHigh(AX)
		Apply4xLowHigh(PADDL)
		Store4xHigh(AX)

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
