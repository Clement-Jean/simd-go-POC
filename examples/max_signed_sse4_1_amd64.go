package main

// Max8x16 is available only starting from SSE4.1
// As such, this file will only compile with the
// `go build -flags sse4.1 ...` command

import (
	"fmt"
	"simd"
)

func main() {
	a := [16]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b := [16]int8{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("%v\n", simd.Max8x16(a, b))
}
