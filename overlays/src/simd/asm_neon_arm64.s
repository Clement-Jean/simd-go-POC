//go:build neon

#include "textflag.h"

TEXT ·maxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MaxU8x16(SB)

TEXT ·max8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Max8x16(SB)

TEXT ·minU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·MinU8x16(SB)

TEXT ·min8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Min8x16(SB)

TEXT ·reduceMaxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMaxU8x16(SB)

TEXT ·reduceMax8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMax8x16(SB)

TEXT ·reduceMinU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMinU8x16(SB)

TEXT ·reduceMin8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMin8x16(SB)

TEXT ·extract8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·Extract8x16(SB)

TEXT ·extractU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ExtractU8x16(SB)
