package simd

func And8x16(out, a, b *[16]uint8)
func And16x8(out, a, b *[8]uint16)
func Or8x16(out, a, b *[16]uint8)
func Xor8x16(out, a, b *[16]uint8)
func ShiftLeft8x16(out, a *[16]uint8, n uint)
func ShiftRight8x16(out, a *[16]uint8, n uint)
func ShiftRight16x8(out, a *Uint16x8, n uint)
func AllZerosU8x16(a *[16]uint8) bool
