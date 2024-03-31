#include "textflag.h"

// func Max8x16(out, a, b *[16]int8)
TEXT ·Max8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x4e216400
	VST1 [V0.D2], (R0)
	RET

// func MaxU8x16(out, a, b *[16]uint8)
TEXT ·MaxU8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x6e216400
	VST1 [V0.D2], (R0)
	RET

// func Min8x16(out, a, b *[16]int8)
TEXT ·Min8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x4e216c00
	VST1 [V0.D2], (R0)
	RET

// func MinU8x16(out, a, b *[16]uint8)
TEXT ·MinU8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x6e216c00
	VST1 [V0.D2], (R0)
	RET
