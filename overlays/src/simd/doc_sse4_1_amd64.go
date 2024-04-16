//go:build sse4.1

package simd

func maxU8x16(out, a, b *[16]uint8)

func max8x16(out, a, b *[16]int8)

func MaxU8x16(a, b [16]uint8) (out [16]uint8) {
	maxU8x16(&out, &a, &b)
	return out
}

func Max8x16(a, b [16]int8) (out [16]int8) {
	max8x16(&out, &a, &b)
	return out
}

func minU8x16(out, a, b *[16]uint8)

func min8x16(out, a, b *[16]int8)

func MinU8x16(a, b [16]uint8) (out [16]uint8) {
	minU8x16(&out, &a, &b)
	return out
}

func Min8x16(a, b [16]int8) (out [16]int8) {
	min8x16(&out, &a, &b)
	return out
}
