//go:build sse4.1

package simd

func max8x16(out, a, b *Int8x16)

func Max8x16(a, b Int8x16) (out Int8x16) {
	max8x16(&out, &a, &b)
	return out
}

func min8x16(out, a, b *Int8x16)

func Min8x16(a, b Int8x16) (out Int8x16) {
	min8x16(&out, &a, &b)
	return out
}

func allZerosU8x16(a *Uint8x16) bool
func allZeros8x16(a *Int8x16) bool

func AllZerosU8x16(a Uint8x16) bool {
	return allZerosU8x16(&a)
}

func AllZeros8x16(a Int8x16) bool {
	return allZeros8x16(&a)
}
