#include "go_asm.h"
#include "textflag.h"
#include "funcdata.h"

// func qpc() int64
TEXT main·asmQPC(SB),NOSPLIT,$0-8
        LEAQ    ret+0(FP), CX
        MOVQ    SP, AX
        ANDQ    $~15, SP        // alignment as per Windows requirement
        SUBQ    $(48), SP       // room for SP and 4 args as per Windows require
                                // plus one extra word to keep stack 16 bytes al
        MOVQ    AX, 32(SP)
        MOVQ    main·_QueryPerformanceCounter(SB), AX
        CALL    AX
        MOVQ    32(SP), SP
        RET
