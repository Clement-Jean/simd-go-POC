//go:build sse2 || sse3 || sse4.1

package simd

func MoveByteMaskU8x16(a *[16]uint8) uint16
