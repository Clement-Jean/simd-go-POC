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
