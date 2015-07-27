#include "textflag.h"

// func AddU32_ASM(dst, src []uint32)
TEXT ·AddU32_ASM(SB),NOSPLIT,$0-48
	MOVQ     dst+8(FP),  CX // CX = len(dst)
	MOVQ     src+32(FP), BX // BX = len(src)

	CMPQ     CX, BX         //
	CMOVQGT  BX, CX         // CX = min(BX, CX)

	MOVQ     dst+0(FP),  AX // AX = &dst[0]
	MOVQ     src+24(FP), BX // BX = &src[0]

	INCQ     CX             // CX++
next:
	DECQ     CX             // CX--
	JZ done                 // jump if CX = 0

	MOVL (BX),  DX          //
	ADDL   DX, (AX)         // [AX] += [BX]

	ADDQ $4, BX             // BX += 4
	ADDQ $4, AX             // AX += 4

	JMP next

done:
	RET

// func SubU32_ASM(dst, src []uint32)
TEXT ·SubU32_ASM(SB),NOSPLIT,$0
	MOVQ     dst+8(FP),  CX // CX = len(dst)
	MOVQ     src+32(FP), BX // BX = len(src)

	CMPQ     CX, BX         //
	CMOVQGT  BX, CX         // CX = min(BX, CX)

	MOVQ     dst+0(FP),  AX // AX = &dst[0]
	MOVQ     src+24(FP), BX // BX = &src[0]

	INCQ     CX             // CX++
next:
	DECQ     CX             // CX--
	JZ done                 // jump if CX = 0

	MOVL     (BX), DX       // DX = [BX]
	SUBL     DX, (AX)       // [AX] -= DX

	ADDQ  $4, BX            // BX += 4
	ADDQ  $4, AX            // AX += 4

	JMP next

done:
    RET

// func MulU32_ASM(dst, src []uint32)
TEXT ·MulU32_ASM(SB),NOSPLIT,$0
	MOVQ     dst+8(FP),  CX // CX = len(dst)
	MOVQ     src+32(FP), BX // BX = len(src)

	CMPQ     CX, BX         //
	CMOVQGT  BX, CX         // CX = min(BX, CX)

	MOVQ     dst+0(FP),  AX // AX = &dst[0]
	MOVQ     src+24(FP), BX // BX = &src[0]

	INCQ     CX             // CX++
next:
	DECQ     CX             // CX--
	JZ done                 // jump if CX = 0

	MOVL  (BX), DX          // DX = [BX]
	IMULL (AX), DX          // DX *= [AX]
	MOVL  DX, (AX)          // [AX] = DX

	ADDQ $4, BX             // BX += 4
	ADDQ $4, AX             // AX += 4

	JMP next

done:
    RET

