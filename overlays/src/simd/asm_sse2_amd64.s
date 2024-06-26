//go:build sse2 || sse3 || sse4.1

#include "textflag.h"

TEXT ·addU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Add8x16(SB)

TEXT ·add8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Add8x16(SB)

TEXT ·saturatingAddU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·SaturatingAddU8x16(SB)

TEXT ·saturatingAdd8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·SaturatingAdd8x16(SB)

TEXT ·subU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Sub8x16(SB)

TEXT ·sub8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Sub8x16(SB)

TEXT ·saturatingSubU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·SaturatingSubU8x16(SB)

TEXT ·saturatingSub8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·SaturatingSub8x16(SB)

TEXT ·andU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·And8x16(SB)

TEXT ·and8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·And8x16(SB)

TEXT ·andU16x8(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·And16x8(SB)

TEXT ·and16x8(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·And16x8(SB)

TEXT ·orU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Or8x16(SB)

TEXT ·or8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Or8x16(SB)

TEXT ·xorU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Xor8x16(SB)

TEXT ·xor8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Xor8x16(SB)

TEXT ·maxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MaxU8x16(SB)

TEXT ·minU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MinU8x16(SB)

TEXT ·moveByteMaskU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MoveByteMaskU8x16(SB)

TEXT ·moveByteMask8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MoveByteMaskU8x16(SB)
