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
