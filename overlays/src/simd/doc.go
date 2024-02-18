package simd

func addU8x16(out, a, b *[16]uint8)

func add8x16(out, a, b *[16]int8)

func AddU8x16(a, b *[16]uint8) *[16]uint8 {
	out := new([16]uint8)
	//fmt.Printf("a -> %+v\n", a)
	//fmt.Printf("b -> %+v\n", b)
	//fmt.Printf("out -> %+v\n", out)
	//fmt.Printf("------------\n")
	addU8x16(out, a, b)
	//fmt.Printf("a -> %+v\n", a)
	//fmt.Printf("b -> %+v\n", b)
	//fmt.Printf("out -> %+v\n", out)
	//fmt.Printf("------------\n")
	return out
}

func Add8x16(a, b *[16]int8) *[16]int8 {
	out := new([16]int8)
	// fmt.Printf("a -> %+v\n", a)
	// fmt.Printf("b -> %+v\n", b)
	// fmt.Printf("out -> %+v\n", out)
	// fmt.Printf("------------\n")
	add8x16(out, a, b)
	// fmt.Printf("a -> %+v\n", a)
	// fmt.Printf("b -> %+v\n", b)
	// fmt.Printf("out -> %+v\n", out)
	// fmt.Printf("------------\n")
	return out
}
