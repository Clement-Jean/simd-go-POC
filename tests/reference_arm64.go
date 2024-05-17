package tests

import (
	"unsafe"

	"example.com/tests/internal"
)

func referenceAdd8x16(a, b [16]int8) (result [16]int8) {
	internal.VaddqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceAddU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VaddqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingAdd8x16(a, b [16]int8) (result [16]int8) {
	internal.VqaddqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingAddU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VqaddqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSub8x16(a, b [16]int8) (result [16]int8) {
	internal.VsubqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSubU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VsubqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingSub8x16(a, b [16]int8) (result [16]int8) {
	internal.VqsubqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceSaturatingSubU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VqsubqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceAnd8x16(a, b [16]int8) (result [16]int8) {
	internal.VandqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceAndU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VandqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceOr8x16(a, b [16]int8) (result [16]int8) {
	internal.VorrqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceOrU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VorrqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceXor8x16(a, b [16]int8) (result [16]int8) {
	internal.VeorqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceXorU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VeorqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceShiftRight8x16(a [16]int8, n uint) (result [16]int8) {
	// the current implementation of ShiftRight is the logical shift right
	// this means that the underlying C intrinsic is vshrq_n_u8 and not
	// vshrq_n_s8
	internal.VshrqnU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		n,
	)
	return result
}

func referenceShiftRightU8x16(a [16]uint8, n uint) (result [16]uint8) {
	internal.VshrqnU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		n,
	)
	return result
}

func referenceMax8x16(a, b [16]int8) (result [16]int8) {
	internal.VmaxqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMaxU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VmaxqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMin8x16(a, b [16]int8) (result [16]int8) {
	internal.VminqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceMinU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.VminqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceReduceMax8x16(a [16]int8) (result int8) {
	internal.VmaxvqS8(
		(*internal.Int8)(&result),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
	)
	return result
}

func referenceReduceMaxU8x16(a [16]uint8) (result uint8) {
	internal.VmaxvqU8(
		(*internal.Uint8)(&result),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
	)
	return result
}

func referenceReduceMin8x16(a [16]int8) (result int8) {
	internal.VminvqS8(
		(*internal.Int8)(&result),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
	)
	return result
}

func referenceReduceMinU8x16(a [16]uint8) (result uint8) {
	internal.VminvqU8(
		(*internal.Uint8)(&result),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
	)
	return result
}

func referenceExtract8x16(a, b [16]int8, n int) (result [16]int8) {
	internal.VextqS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Int8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceExtractU8x16(a, b [16]uint8, n int) (result [16]uint8) {
	internal.VextqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceLookup8x16(a [16]int8, b [16]uint8) (result [16]int8) {
	internal.Vqtbl1qS8(
		(*internal.Int8x16)(unsafe.Pointer(&result)),
		(*internal.Int8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}

func referenceLookupU8x16(a, b [16]uint8) (result [16]uint8) {
	internal.Vqtbl1qU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result)),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)
	return result
}
