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
