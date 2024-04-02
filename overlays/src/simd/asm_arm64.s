#include "textflag.h"
	
TEXT ·ReduceMaxU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMaxU8x16(SB)

TEXT ·ReduceMax8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMax8x16(SB)

TEXT ·ReduceMinU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMinU8x16(SB)

TEXT ·ReduceMin8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·ReduceMin8x16(SB)
