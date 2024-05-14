//go:build sse4.1

package simd

func max8x16(out, a, b *[16]int8)

func Max8x16(a, b [16]int8) (out [16]int8) {
	max8x16(&out, &a, &b)
	return out
}

func min8x16(out, a, b *[16]int8)

func Min8x16(a, b [16]int8) (out [16]int8) {
	min8x16(&out, &a, &b)
	return out
}

func allZerosU8x16(a *[16]uint8) bool
func allZeros8x16(a *[16]int8) bool

func AllZerosU8x16(a [16]uint8) bool {
	return allZerosU8x16(&a)
}

func AllZeros8x16(a [16]int8) bool {
	return allZeros8x16(&a)
}
