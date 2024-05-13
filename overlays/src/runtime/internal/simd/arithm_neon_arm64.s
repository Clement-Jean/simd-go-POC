#include "textflag.h"

// func Add8x16(out, a, b *[16]uint8)
TEXT ·Add8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	VADD V1.B16, V0.B16, V0.B16
	VST1 [V0.D2], (R0)
	RET

// func SaturatingAddU8x16(out, a, b *[16]uint8)
TEXT ·SaturatingAddU8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x6e210c00
	VST1 [V0.D2], (R0)
	RET

// func SaturatingAdd8x16(out, a, b *[16]int8)
TEXT ·SaturatingAdd8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x4e210c00
	VST1 [V0.D2], (R0)
	RET

// func Sub8x16(out, a, b *[16]uint8)
TEXT ·Sub8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	VSUB V1.B16, V0.B16, V0.B16
	VST1 [V0.D2], (R0)
	RET

// func SaturatingSubU8x16(out, a, b *[16]uint8)
TEXT ·SaturatingSubU8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x6e212c00
	VST1 [V0.D2], (R0)
	RET

// func SaturatingSub8x16(out, a, b *[16]int8)
TEXT ·SaturatingSub8x16(SB),NOSPLIT,$0-24
	MOVD out+0(FP), R0
	MOVD a+8(FP), R1
	MOVD b+16(FP), R2
	VLD1 (R1), [V0.D2]
	VLD1 (R2), [V1.D2]
	WORD $0x4e212c00
	VST1 [V0.D2], (R0)
	RET
