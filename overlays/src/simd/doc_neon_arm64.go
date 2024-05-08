//go:build neon

package simd

func splatU8x16(out *Uint8x16, n uint8)

func splat8x16(out *Int8x16, n int8)

func SplatU8x16(n uint8) (out Uint8x16) {
	splatU8x16(&out, n)
	return out
}

func Splat8x16(n int8) (out Int8x16) {
	splat8x16(&out, n)
	return out
}

func maxU8x16(out, a, b *Uint8x16)

func max8x16(out, a, b *Int8x16)

func MaxU8x16(a, b Uint8x16) (out Uint8x16) {
	maxU8x16(&out, &a, &b)
	return out
}

func Max8x16(a, b Int8x16) (out Int8x16) {
	max8x16(&out, &a, &b)
	return out
}

func minU8x16(out, a, b *Uint8x16)

func min8x16(out, a, b *Int8x16)

func MinU8x16(a, b Uint8x16) (out Uint8x16) {
	minU8x16(&out, &a, &b)
	return out
}

func Min8x16(a, b Int8x16) (out Int8x16) {
	min8x16(&out, &a, &b)
	return out
}

func reduceMaxU8x16(a *Uint8x16) uint8
func reduceMax8x16(a *Int8x16) int8

func ReduceMaxU8x16(a Uint8x16) uint8 {
	return reduceMaxU8x16(&a)
}

func ReduceMax8x16(a Int8x16) int8 {
	return reduceMax8x16(&a)
}

func reduceMinU8x16(a *Uint8x16) uint8
func reduceMin8x16(a *Int8x16) int8

func ReduceMinU8x16(a Uint8x16) uint8 {
	return reduceMinU8x16(&a)
}

func ReduceMin8x16(a Int8x16) int8 {
	return reduceMin8x16(&a)
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

func shiftLeft8x16(out, a *Int8x16, idx uint)
func shiftLeftU8x16(out, a *Uint8x16, idx uint)

func ShiftLeft8x16(a Int8x16, idx uint) (out Int8x16) {
	// if idx > 7 {
	// 	return out
	// }
	shiftLeft8x16(&out, &a, idx)
	return out
}

func ShiftLeftU8x16(a Uint8x16, idx uint) (out Uint8x16) {
	// if idx > 7 {
	// 	return out
	// }
	shiftLeftU8x16(&out, &a, idx)
	return out
}

func shiftRight8x16(out, a *Int8x16, idx uint)
func shiftRightU8x16(out, a *Uint8x16, idx uint)

func ShiftRight8x16(a Int8x16, idx uint) (out Int8x16) {
	// if idx < 1 || idx > 8 {
	// 	return out
	// }
	shiftRight8x16(&out, &a, idx)
	return out
}

func ShiftRightU8x16(a Uint8x16, idx uint) (out Uint8x16) {
	// if idx < 1 || idx > 8 {
	// 	return out
	// }
	shiftRightU8x16(&out, &a, idx)
	return out
}
