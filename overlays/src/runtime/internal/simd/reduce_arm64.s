#include "textflag.h"

// func ReduceMax8x16(a *[16]uint8) uint8
TEXT ·ReduceMax8x16(SB),NOSPLIT,$0-9
	MOVD a+0(FP), R0
	VLD1 (R0), [V0.D2]
	WORD $0x6e30a801
	VMOV V1.B[0], R1
	MOVB R1, ret+8(FP)
	RET

// func ReduceMin8x16(a *[16]uint8) uint8
TEXT ·ReduceMin8x16(SB),NOSPLIT,$0-9
	MOVD a+0(FP), R0
	VLD1 (R0), [V0.D2]
	WORD $0x6e31a801
	VMOV V1.B[0], R1
	MOVB R1, ret+8(FP)
	RET
