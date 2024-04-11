package main

// ReduceMaxU8x16 is only available in ARM64 NEON

import (
	"fmt"
	"simd"
)

func main() {
	a := [16]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("%d\n", simd.ReduceMaxU8x16(a))
}
