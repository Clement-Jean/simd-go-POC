package simd

//go:noescape
func Add8x16(out, a, b *[16]uint8)
