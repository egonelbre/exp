#include "textflag.h"
#include "go_asm.h"

// func rdtscp_pid() int
TEXT ·rdtscp_pid(SB), NOSPLIT, $0-8
	BYTE $0x0F; BYTE $0x01; BYTE $0xF9 // RDTSCP
	ANDQ $0xff, CX
	MOVQ CX, ret+0(FP)
	RET

// func rdpid_pid() int
TEXT ·rdpid_pid(SB), NOSPLIT, $0-8
	BYTE $0xF3; BYTE $0x0F; BYTE $0xC7; BYTE $0xf8; // RDPID AX
	MOVL AX, ret+0(FP)
	RET

// func cpuid_pid() int
TEXT ·cpuid_pid(SB), NOSPLIT, $0-8
	MOVL $0xB, AX
	CPUID // .0BH:EDX[31:0]
	MOVL BX, AX
	RET
