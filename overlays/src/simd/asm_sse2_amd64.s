//go:build sse2

#include "textflag.h"

TEXT ·maxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MaxU8x16(SB)

TEXT ·minU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MinU8x16(SB)

TEXT ·movMaskByteU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MovMaskByteU8x16(SB)

TEXT ·movMaskByte8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MoveMaskByteU8x16(SB)
