//go:build ssse3

package simd

func lookup8x16(out, a *Int8x16, b *Uint8x16)
func lookupU8x16(out, a, b *Uint8x16)

func Lookup8x16(a Int8x16, b Uint8x16) (out Int8x16) {
	lookup8x16(&out, &a, &b)
	return out
}

func LookupU8x16(a, b Uint8x16) (out Uint8x16) {
	lookupU8x16(&out, &a, &b)
	return out
}
