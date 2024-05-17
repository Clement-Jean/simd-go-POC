//go:build ssse3

package simd

func extract8x16(out, a, b *Int8x16, idx uint)
func extractU8x16(out, a, b *Uint8x16, idx uint)

func Extract8x16(a, b Int8x16, idx uint) (out Int8x16) {
	extract8x16(&out, &a, &b, idx)
	return out
}

func ExtractU8x16(a, b Uint8x16, idx uint) (out Uint8x16) {
	extractU8x16(&out, &a, &b, idx)
	return out
}

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
