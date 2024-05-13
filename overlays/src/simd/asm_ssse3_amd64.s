//go:build ssse3

#include "textflag.h"

TEXT ·extract8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·Extract8x16(SB)

TEXT ·extractU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ExtractU8x16(SB)

TEXT ·lookup8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·Lookup8x16(SB)

TEXT ·lookupU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·LookupU8x16(SB)
