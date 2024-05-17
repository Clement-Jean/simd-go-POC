//go:build sse4.1

package simd

func Max8x16(out, a, b *[16]int8)
func Min8x16(out, a, b *[16]int8)
