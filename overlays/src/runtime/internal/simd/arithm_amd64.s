#include "textflag.h"

// func Add8x16(out, a, b *[16]uint8)
TEXT ·Add8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVUPS (BX), X0
	VMOVUPS (CX), X1
	VPADDQ X1, X0, X0
	MOVUPD X0, (AX)
	RET

// func SaturatingAddU8x16(out, a, b *[16]uint8)
TEXT ·SaturatingAddU8x16(SB),NOSPLIT,$0-24
	RET

// func SaturatingAdd8x16(out, a, b *[16]int8)
TEXT ·SaturatingAdd8x16(SB),NOSPLIT,$0-24
	RET

// func Sub8x16(out, a, b *[16]uint8)
TEXT ·Sub8x16(SB),NOSPLIT,$0-24
	RET

// func SaturatingSubU8x16(out, a, b *[16]uint8)
TEXT ·SaturatingSubU8x16(SB),NOSPLIT,$0-24
	RET

// func SaturatingSub8x16(out, a, b *[16]int8)
TEXT ·SaturatingSub8x16(SB),NOSPLIT,$0-24
	RET
