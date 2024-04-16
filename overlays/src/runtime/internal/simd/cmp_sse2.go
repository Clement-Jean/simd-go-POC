//go:build sse2

package simd

func MaxU8x16(out, a, b *[16]uint8)
func MinU8x16(out, a, b *[16]uint8)
