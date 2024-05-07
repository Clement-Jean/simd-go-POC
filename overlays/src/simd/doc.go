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

func shiftRight8x16(out, a *Int8x16, idx uint)
func shiftRightU8x16(out, a *Uint8x16, idx uint)

func ShiftRight8x16(a Int8x16, idx uint) (out Int8x16) {
	// if idx < 1 || idx > 8 {
	// 	return out
	// }
	shiftRight8x16(&out, &a, idx)
	return out
}

func ShiftRightU8x16(a Uint8x16, idx uint) (out Uint8x16) {
	// if idx < 1 || idx > 8 {
	// 	return out
	// }
	shiftRightU8x16(&out, &a, idx)
	return out
}

func extract8x16(out, a, b *Int8x16, idx uint)
func extractU8x16(out, a, b *Uint8x16, idx uint)

func Extract8x16(a, b Int8x16, idx uint) (out Int8x16) {
	extract8x16(&out, &a, &b, idx)
	return out
}

func ExtractU8x16(a, b Uint8x16, idx uint) (out Uint8x16) {
	extractU8x16(&out, &a, &b, idx)
	return out
}
