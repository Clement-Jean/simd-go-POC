#include "textflag.h"

TEXT ·addU8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Add8x16(SB)

TEXT ·add8x16(SB),NOSPLIT,$0
	JMP	runtime∕internal∕simd·Add8x16(SB)

// TEXT ·addU8x32(SB),NOSPLIT,$0
// 	JMP	runtime∕internal∕simd·Add8x32(SB)

// TEXT ·add8x32(SB),NOSPLIT,$0
// 	JMP	runtime∕internal∕simd·Add8x32(SB)
