//go:build neon

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

TEXT ·shiftRightU16x8(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ShiftRight16x8(SB)

TEXT ·shiftRight16x8(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ShiftRight16x8(SB)

TEXT ·splatU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·Splat8x16(SB)

TEXT ·splat8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·Splat8x16(SB)

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

TEXT ·lookup8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·Lookup8x16(SB)

TEXT ·lookupU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·LookupU8x16(SB)

TEXT ·shiftLeftU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ShiftLeft8x16(SB)

TEXT ·shiftLeft8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ShiftLeft8x16(SB)

TEXT ·shiftRightU8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ShiftRight8x16(SB)

TEXT ·shiftRight8x16(SB),NOSPLIT,$0
	JMP runtime∕internal∕simd·ShiftRight8x16(SB)

