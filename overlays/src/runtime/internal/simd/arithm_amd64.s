#include "textflag.h"

// func Add8x16(out, a, b *[16]uint8)
TEXT ·Add8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVDQA (BX), X0
	VMOVDQA (CX), X1
	VPADDB X1, X0, X0
	VMOVDQA X0, (AX)
	RET

// func SaturatingAddU8x16(out, a, b *[16]uint8)
TEXT ·SaturatingAddU8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVDQA (BX), X0
	VMOVDQA (CX), X1
	VPADDUSB X1, X0, X0
	VMOVDQA X0, (AX)
	RET

// func SaturatingAdd8x16(out, a, b *[16]int8)
TEXT ·SaturatingAdd8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVDQA (BX), X0
	VMOVDQA (CX), X1
	VPADDSB X1, X0, X0
	VMOVDQA X0, (AX)
	RET

// func Sub8x16(out, a, b *[16]uint8)
TEXT ·Sub8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVDQA (BX), X0
	VMOVDQA (CX), X1
	VPSUBB X1, X0, X0
	VMOVDQA X0, (AX)
	RET

// func SaturatingSubU8x16(out, a, b *[16]uint8)
TEXT ·SaturatingSubU8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVDQA (BX), X0
	VMOVDQA (CX), X1
	VPSUBUSB X1, X0, X0
	VMOVDQA X0, (AX)
	RET

// func SaturatingSub8x16(out, a, b *[16]int8)
TEXT ·SaturatingSub8x16(SB),NOSPLIT,$0-24
	MOVQ out+0(FP), AX
	MOVQ a+8(FP), BX
	MOVQ b+16(FP), CX
	VMOVDQA (BX), X0
	VMOVDQA (CX), X1
	VPSUBSB X1, X0, X0
	VMOVDQA X0, (AX)
	RET
