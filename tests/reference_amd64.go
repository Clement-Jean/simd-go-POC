package tests

import (
	"unsafe"

	"example.com/tests/internal"
)

func referenceAdd8x16(a, b *[16]uint8) *[16]uint8 {
	result := new([16]int8)
	internal.MmAddEpi8(
		(*internal.M128I)(unsafe.Pointer(result)),
		(*internal.M128I)(unsafe.Pointer(a)),
		(*internal.M128I)(unsafe.Pointer(b)),
	)
	return result
}
