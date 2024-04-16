//go:build sse2

package simd

func maxU8x16(out, a, b *[16]uint8)

func MaxU8x16(a, b [16]uint8) (out [16]uint8) {
	maxU8x16(&out, &a, &b)
	return out
}

func minU8x16(out, a, b *[16]uint8)

func MinU8x16(a, b [16]uint8) (out [16]uint8) {
	minU8x16(&out, &a, &b)
	return out
}
