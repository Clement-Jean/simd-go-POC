package tests

import (
	"unsafe"

	"example.com/tests/internal"
)

func referenceAdd8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VaddqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceAddU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VaddqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingAdd8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VqaddqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingAddU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VqaddqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSub8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VsubqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSubU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VsubqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingSub8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VqsubqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceSaturatingSubU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VqsubqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceAnd8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VandqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceAndU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VandqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceOr8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VorrqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceOrU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VorrqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceXor8x16(a, b *[16]int8) *[16]int8 {
	result := new([16]int8)
	internal.VeorqS8(
		(*internal.Int8x16)(unsafe.Pointer(result)),
		(*internal.Int8x16)(unsafe.Pointer(a)),
		(*internal.Int8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceXorU8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]uint8)
	internal.VeorqU8(
		(*internal.Uint8x16)(unsafe.Pointer(result)),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
		(*internal.Uint8x16)(unsafe.Pointer(b)),
	)
	return result
}

func referenceMax8x16(a *[16]int8) int8 {
	var result int8
	internal.VmaxvqS8(
		(*internal.Int8)(&result),
		(*internal.Int8x16)(unsafe.Pointer(a)),
	)
	return result
}

func referenceMaxU8x16(a *[16]uint8) uint8 {
	var result uint8
	internal.VmaxvqU8(
		(*internal.Uint8)(&result),
		(*internal.Uint8x16)(unsafe.Pointer(a)),
	)
	return result
}
