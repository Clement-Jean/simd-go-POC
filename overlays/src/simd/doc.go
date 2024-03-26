package simd

func addU8x16(out, a, b *[16]uint8)

func add8x16(out, a, b *[16]int8)

func AddU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	addU8x16(out, a, b)
	return out
}

func Add8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	add8x16(out, a, b)
	return out
}

func saturatingAddU8x16(out, a, b *[16]uint8)

func saturatingAdd8x16(out, a, b *[16]int8)

func SaturatingAddU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	saturatingAddU8x16(out, a, b)
	return out
}

func SaturatingAdd8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	saturatingAdd8x16(out, a, b)
	return out
}

func subU8x16(out, a, b *[16]uint8)

func sub8x16(out, a, b *[16]int8)

func SubU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	subU8x16(out, a, b)
	return out
}

func Sub8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	sub8x16(out, a, b)
	return out
}

func saturatingSubU8x16(out, a, b *[16]uint8)

func saturatingSub8x16(out, a, b *[16]int8)

func SaturatingSubU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	saturatingSubU8x16(out, a, b)
	return out
}

func SaturatingSub8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	saturatingSub8x16(out, a, b)
	return out
}

func andU8x16(out, a, b *[16]uint8)

func and8x16(out, a, b *[16]int8)

func AndU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	andU8x16(out, a, b)
	return out
}

func And8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	and8x16(out, a, b)
	return out
}

func orU8x16(out, a, b *[16]uint8)

func or8x16(out, a, b *[16]int8)

func OrU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	orU8x16(out, a, b)
	return out
}

func Or8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	or8x16(out, a, b)
	return out
}

func xorU8x16(out, a, b *[16]uint8)

func xor8x16(out, a, b *[16]int8)

func XorU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	xorU8x16(out, a, b)
	return out
}

func Xor8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	xor8x16(out, a, b)
	return out
}

func MaxU8x16(a *[16]uint8) uint8
func Max8x16(a *[16]int8) int8

func MinU8x16(a *[16]uint8) uint8
func Min8x16(a *[16]int8) int8
