//go:build sse4.1

#include "textflag.h"

TEXT ·maxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MaxU8x16(SB)

TEXT ·max8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Max8x16(SB)

TEXT ·minU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MinU8x16(SB)

TEXT ·min8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Min8x16(SB)
