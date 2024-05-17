//go:build sse4.1

#include "textflag.h"

TEXT ·allZerosU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·AllZerosU8x16(SB)
