package main

import (
	"fmt"
	"simd"
)

func main() {
	a := &[16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b := &[16]uint8{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	c := simd.AddU8x16(a, b)

	fmt.Printf("%v\n", c)

	b = &[16]uint8{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	c = simd.SaturatingAddU8x16(a, b)

	fmt.Printf("%v\n", c)

	c = simd.AddU8x16(a, b)

	fmt.Printf("%v\n", c)

	d := &[16]int8{127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127}
	e := &[16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	f := simd.SaturatingAdd8x16(d, e)

	fmt.Printf("%v\n", f)

	d = &[16]int8{-128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128}
	e = &[16]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	f = simd.SaturatingAdd8x16(d, e)

	fmt.Printf("%v\n", f)

	g := simd.SaturatingSub8x16(d, e)

	fmt.Printf("%v\n", g)

	b = &[16]uint8{}
	h := simd.SaturatingSubU8x16(b, a)

	fmt.Printf("%v\n", h)
	// d := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	// e := simd.AddU8x16((*[16]uint8)(d[16:32]), (*[16]uint8)(d[0:16]))

	// fmt.Printf("%v\n", e)

	// f := &[16]int8{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16}
	// g := &[16]int8{-17, -18, -19, -20, -21, -22, -23, -24, -25, -26, -27, -28, -29, -30, -31, -32}
	// h := simd.Add8x16(f, g)

	// fmt.Printf("%v\n", h)

	// i := []int8{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20, -21, -22, -23, -24, -25, -26, -27, -28, -29, -30, -31, -32}
	// j := simd.Add8x16((*[16]int8)(i[16:32]), (*[16]int8)(i[0:16]))

	// fmt.Printf("%v\n", j)
}
