//go:build sse2 || sse3 || sse4.1

package simd

func addU8x16(out, a, b *Uint8x16)
func add8x16(out, a, b *Int8x16)

func AddU8x16(a, b Uint8x16) (out Uint8x16) {
	addU8x16(&out, &a, &b)
	return out
}

func Add8x16(a, b Int8x16) (out Int8x16) {
	add8x16(&out, &a, &b)
	return out
}

func saturatingAddU8x16(out, a, b *Uint8x16)
func saturatingAdd8x16(out, a, b *Int8x16)

func SaturatingAddU8x16(a, b Uint8x16) (out Uint8x16) {
	saturatingAddU8x16(&out, &a, &b)
	return out
}

func SaturatingAdd8x16(a, b Int8x16) (out Int8x16) {
	saturatingAdd8x16(&out, &a, &b)
	return out
}

func subU8x16(out, a, b *Uint8x16)
func sub8x16(out, a, b *Int8x16)

func SubU8x16(a, b Uint8x16) (out Uint8x16) {
	subU8x16(&out, &a, &b)
	return out
}

func Sub8x16(a, b Int8x16) (out Int8x16) {
	sub8x16(&out, &a, &b)
	return out
}

func saturatingSubU8x16(out, a, b *Uint8x16)
func saturatingSub8x16(out, a, b *Int8x16)

func SaturatingSubU8x16(a, b Uint8x16) (out Uint8x16) {
	saturatingSubU8x16(&out, &a, &b)
	return out
}

func SaturatingSub8x16(a, b Int8x16) (out Int8x16) {
	saturatingSub8x16(&out, &a, &b)
	return out
}

func andU8x16(out, a, b *Uint8x16)
func and8x16(out, a, b *Int8x16)

func AndU8x16(a, b Uint8x16) (out Uint8x16) {
	andU8x16(&out, &a, &b)
	return out
}

func And8x16(a, b Int8x16) (out Int8x16) {
	and8x16(&out, &a, &b)
	return out
}

func andU16x8(out, a, b *Uint16x8)
func and16x8(out, a, b *Int16x8)

func AndU16x8(a, b Uint16x8) (out Uint16x8) {
	andU16x8(&out, &a, &b)
	return out
}

func And16x8(a, b Int16x8) (out Int16x8) {
	and16x8(&out, &a, &b)
	return out
}

func orU8x16(out, a, b *Uint8x16)
func or8x16(out, a, b *Int8x16)

func OrU8x16(a, b Uint8x16) (out Uint8x16) {
	orU8x16(&out, &a, &b)
	return out
}

func Or8x16(a, b Int8x16) (out Int8x16) {
	or8x16(&out, &a, &b)
	return out
}

func xorU8x16(out, a, b *Uint8x16)
func xor8x16(out, a, b *Int8x16)

func XorU8x16(a, b Uint8x16) (out Uint8x16) {
	xorU8x16(&out, &a, &b)
	return out
}

func Xor8x16(a, b Int8x16) (out Int8x16) {
	xor8x16(&out, &a, &b)
	return out
}

func shiftRight16x8(out, a *Uint16x8, idx uint)
func shiftRightU16x8(out, a *Uint16x8, idx uint)

func ShiftRight16x8(a Uint16x8, idx uint) (out Uint16x8) {
	shiftRight16x8(&out, &a, idx)
	return out
}

func ShiftRightU16x8(a Uint16x8, idx uint) (out Uint16x8) {
	shiftRightU16x8(&out, &a, idx)
	return out
}

func maxU8x16(out, a, b *Uint8x16)

func MaxU8x16(a, b Uint8x16) (out Uint8x16) {
	maxU8x16(&out, &a, &b)
	return out
}

func minU8x16(out, a, b *Uint8x16)

func MinU8x16(a, b Uint8x16) (out Uint8x16) {
	minU8x16(&out, &a, &b)
	return out
}

func movMaskByteU8x16(a *Uint8x16) uint16
func movMaskByte8x16(a *Int8x16) uint16

func MovMaskByteU8x16(a Uint8x16) uint16 {
	return movMaskByteU8x16(&a)
}

func MovMaskByte8x16(a Int8x16) uint16 {
	return movMaskByte8x16(&a)
}
