package tests

import (
	"unsafe"

	"example.com/tests/internal"
)

func referenceAdd8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmAddEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAddU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmAddEpi8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSaturatingAdd8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmAddsEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSaturatingAddU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmAddsEpu8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSub8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmSubEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSubU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmSubEpi8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSaturatingSub8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmSubsEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSaturatingSubU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmSubsEpu8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAnd8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmAndSi128(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAndU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmAndSi128(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAnd16x8(a, b [8]int16) (result [8]int16) {
	_a := internal.MmSetrEpi16(a)
	_b := internal.MmSetrEpi16(b)
	_result := internal.MmAndSi128(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAndU16x8(a, b [8]uint16) (result [8]uint16) {
	_a := internal.MmSetrEpu16(a)
	_b := internal.MmSetrEpu16(b)
	_result := internal.MmAndSi128(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceOr8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmOrSi128(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceOrU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmOrSi128(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceXor8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmXorSi128(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceXorU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmXorSi128(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceMax8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmMaxEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceMaxU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmMaxEpu8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceMin8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmMinEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceMinU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmMinEpu8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceExtract8x16(a, b [16]int8, imm8 int) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmAlignrEpi8(_b, _a, imm8)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceExtractU8x16(a, b [16]uint8, imm8 int) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmAlignrEpi8(_b, _a, imm8)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceLookup8x16(a [16]int8, b [16]uint8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmShuffleEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Int8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceLookupU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpu8(a)
	_b := internal.MmSetrEpu8(b)
	_result := internal.MmShuffleEpi8(_a, _b)
	internal.MmStoreuSu128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAllZeros8x16(a [16]int8) bool {
	_a := internal.MmSetrEpi8(a)
	return internal.MmTestAllZeros(_a)
}

func referenceAllZerosU8x16(a [16]uint8) bool {
	_a := internal.MmSetrEpu8(a)
	return internal.MmTestAllZeros(_a)
}
