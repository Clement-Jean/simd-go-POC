//go:build sse4.1

package simd

func allZerosU8x16(a *[16]uint8) bool
func allZeros8x16(a *[16]int8) bool

func AllZerosU8x16(a [16]uint8) bool {
	return allZerosU8x16(&a)
}

func AllZeros8x16(a [16]int8) bool {
	return allZeros8x16(&a)
}
