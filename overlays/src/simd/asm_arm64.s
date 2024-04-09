#include "textflag.h"

TEXT ·reduceMaxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMaxU8x16(SB)

TEXT ·reduceMax8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMax8x16(SB)

TEXT ·reduceMinU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMinU8x16(SB)

TEXT ·reduceMin8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMin8x16(SB)
