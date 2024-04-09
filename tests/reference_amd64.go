package tests

import (
	"unsafe"

	"example.com/tests/internal"
)

func referenceAdd8x16(a, b [16]int8) (result [16]int8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmAddEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceAddU8x16(a, b [16]uint8) (result [16]uint8) {
	_a := internal.MmSetrEpi8(a)
	_b := internal.MmSetrEpi8(b)
	_result := internal.MmAddEpi8(_a, _b)
	internal.MmStoreuSi128((*internal.Uint8)(unsafe.Pointer(&result[0])), _result)
	return result
}

func referenceSaturatingAdd8x16(a, b [16]int8) (result [16]int8) {
	internal.MmAddsEpi8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingAddU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmAddsEpu8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSub8x16(a, b [16]int8) (result [16]int8) {
	internal.MmSubEpi8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSubU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmSubEpi8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingSub8x16(a, b [16]int8) (result [16]int8) {
	internal.MmSubsEpi8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingSubU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmSubsEpu8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceAnd8x16(a, b [16]int8) (result [16]int8) {
	internal.MmAndSi128(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceAndU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmAndSi128(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceOr8x16(a, b [16]int8) (result [16]int8) {
	internal.MmOrSi128(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceOrU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmOrSi128(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceXor8x16(a, b [16]int8) (result [16]int8) {
	internal.MmXorSi128(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceXorU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmXorSi128(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMax8x16(a, b [16]int8) (result [16]int8) {
	internal.MmMaxEpi8( // requires SSE4.1
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMaxU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmMaxEpu8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMin8x16(a, b [16]int8) (result [16]int8) {
	internal.MmMinEpi8( // requires SSE4.1
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMinU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.MmMinEpu8(
		(*internal.M128I)(unsafe.Pointer(&result)),
		(*internal.M128I)(unsafe.Pointer(&a)),
		(*internal.M128I)(unsafe.Pointer(&b)),
	)
	return result
}
