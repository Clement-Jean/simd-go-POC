package main

// A generic intrinsic is available in every ISA
// e.g. AddU8x16

import (
	"fmt"
	"simd"
)

func main() {
	a := simd.Uint8x16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b := simd.Uint8x16{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("%v\n", simd.AddU8x16(a, b))
}
