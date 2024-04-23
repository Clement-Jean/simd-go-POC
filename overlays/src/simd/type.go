package simd

import "unsafe"

type Int8x16 [16]int8
type Uint8x16 [16]uint8
type Uint16x8 [8]uint16
type Int16x8 [8]int16
type Uint32x4 [4]uint32
type Int32x4 [4]int32

func AsUint8x16[U Uint16x8 | Uint32x4](u U) Uint8x16 {
	return *(*Uint8x16)(unsafe.Pointer(&u))
}

func AsUint16x8[U Uint8x16 | Uint32x4](u U) Uint16x8 {
	return *(*Uint16x8)(unsafe.Pointer(&u))
}
