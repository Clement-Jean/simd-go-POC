package main

import (
	"simd"
	"unicode/utf8"
)

func validString(in string) bool {
	var prevIncomplete, prevInputBlock simd.Uint8x16
	var processedLen int

	for ; processedLen+16 <= len(in); processedLen += 16 {
		currBlock := *(*simd.Uint8x16)([]byte(in[processedLen:]))

		if simd.MoveByteMaskU8x16(currBlock) < 0x80 {
			if simd.MoveByteMaskU8x16(prevIncomplete) != 0 {
				return false
			}
			prevIncomplete = simd.Uint8x16{}
		} else {
			prev1 := simd.ExtractU8x16(prevInputBlock, currBlock, 15)
			s := simd.AndU16x8(simd.ShiftRightU16x8(simd.AsUint16x8(prev1), 4), ffBy4)
			byte1High := simd.LookupU8x16(shuf1, simd.AsUint8x16(s))
			byte1Low := simd.LookupU8x16(shuf2, simd.AndU8x16(prev1, v0f))
			s = simd.AndU16x8(simd.ShiftRightU16x8(simd.AsUint16x8(currBlock), 4), ffBy4)
			byte2High := simd.LookupU8x16(shuf3, simd.AsUint8x16(s))
			sc := simd.AndU8x16(simd.AndU8x16(byte1High, byte1Low), byte2High)
			prev2 := simd.ExtractU8x16(prevInputBlock, currBlock, 14)
			prev3 := simd.ExtractU8x16(prevInputBlock, currBlock, 13)
			prevInputBlock = currBlock
			isThirdByte := simd.SaturatingSubU8x16(prev2, thirdByte)
			isFourthByte := simd.SaturatingSubU8x16(prev3, fourthByte)
			must23 := simd.OrU8x16(isThirdByte, isFourthByte)
			must23As80 := simd.AndU8x16(must23, v80)
			if err := simd.XorU8x16(must23As80, sc); !simd.AllZerosU8x16(err) {
				return false
			}
			prevIncomplete = simd.SaturatingSubU8x16(currBlock, maxValue)
		}
	}

	if processedLen < len(in) {
		if processedLen > 0 && int8(in[processedLen]) <= -65 {
			processedLen -= 1

			if processedLen >= 0 && int8(in[processedLen]) <= -65 {
				processedLen -= 1
			}
		}
		return utf8.ValidString(in[processedLen:])
	}
	return simd.MoveByteMaskU8x16(prevIncomplete) == 0
}
