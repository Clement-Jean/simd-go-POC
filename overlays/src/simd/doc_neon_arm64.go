//go:build neon

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

func reduceMaxU8x16(a *[16]uint8) uint8
func reduceMax8x16(a *[16]int8) int8

func ReduceMaxU8x16(a [16]uint8) uint8 {
	return reduceMaxU8x16(&a)
}

func ReduceMax8x16(a [16]int8) int8 {
	return reduceMax8x16(&a)
}

func reduceMinU8x16(a *[16]uint8) uint8
func reduceMin8x16(a *[16]int8) int8

func ReduceMinU8x16(a [16]uint8) uint8 {
	return reduceMinU8x16(&a)
}

func ReduceMin8x16(a [16]int8) int8 {
	return reduceMin8x16(&a)
}

func extract8x16(out, a, b *[16]int8, idx uint)
func extractU8x16(out, a, b *[16]uint8, idx uint)

func Extract8x16(a, b [16]int8, idx uint) (out [16]int8) {
	extract8x16(&out, &a, &b, idx)
	return out
}

func ExtractU8x16(a, b [16]uint8, idx uint) (out [16]uint8) {
	extractU8x16(&out, &a, &b, idx)
	return out
}

func lookup8x16(out, a, b *[16]int8)
func lookupU8x16(out, a, b *[16]uint8)

func Lookup8x16(a, b [16]int8) (out [16]int8) {
	lookup8x16(&out, &a, &b)
	return out
}

func LookupU8x16(a, b [16]uint8) (out [16]uint8) {
	lookupU8x16(&out, &a, &b)
	return out
}
