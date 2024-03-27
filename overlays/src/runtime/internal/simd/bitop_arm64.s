#include "textflag.h"

// func And8x16(out, a, b *[16]uint8)
TEXT ·And8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	VAND V1.B16, V0.B16, V0.B16
	VST1 [V0.D2], (R0)
	RET

// func Or8x16(out, a, b *[16]uint8)
TEXT ·Or8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	VORR V1.B16, V0.B16, V0.B16
	VST1 [V0.D2], (R0)
	RET

// func Xor8x16(out, a, b *[16]uint8)
TEXT ·Xor8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	VEOR V1.B16, V0.B16, V0.B16
	VST1 [V0.D2], (R0)
	RET
