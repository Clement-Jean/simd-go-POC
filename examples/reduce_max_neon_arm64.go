package main

// ReduceMaxU8x16 is only available in ARM64 NEON
// As such, this file will only compile with the
// `go build -flags neon ...` command

import (
	"fmt"
	"simd"
)

func main() {
	a := simd.Uint8x16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("%d\n", simd.ReduceMaxU8x16(a))
}
