package tests

import (
	"unsafe"

	"example.com/tests/internal"
)

func referenceAdd8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.MmAddEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceAddU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.MmAddEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingAdd8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.MmAddsEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingAddU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.MmAddsEpu8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSub8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.MmSubEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSubU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.MmSubEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingSub8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.MmSubsEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingSubU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.MmSubsEpu8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}

func referenceAnd8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	return result
}

func referenceAndU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	return result
}

func referenceOr8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	return result
}

func referenceOrU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	return result
}

func referenceXor8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	return result
}

func referenceXorU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	return result
}

func referenceMax8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	return result
}

func referenceMaxU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	return result
}

func referenceMin8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	return result
}

func referenceMinU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	return result
}

func referenceReduceMax8x16(a *[16]int8) int8 {
	var result int8
	return result
}

func referenceReduceMaxU8x16(a *[16]uint8) uint8 {
	var result uint8
	return result
}

func referenceReduceMin8x16(a *[16]int8) int8 {
	var result int8
	return result
}

func referenceReduceMinU8x16(a *[16]uint8) uint8 {
	var result uint8
	return result
}
