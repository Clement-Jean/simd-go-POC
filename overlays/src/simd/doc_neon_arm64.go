//go:build neon

package simd

func addU8x16(out, a, b *Uint8x16)
func add8x16(out, a, b *Int8x16)

func AddU8x16(a, b Uint8x16) (out Uint8x16) {
	addU8x16(&out, &a, &b)
	return out
}

func Add8x16(a, b Int8x16) (out Int8x16) {
	add8x16(&out, &a, &b)
	return out
}

func saturatingAddU8x16(out, a, b *Uint8x16)
func saturatingAdd8x16(out, a, b *Int8x16)

func SaturatingAddU8x16(a, b Uint8x16) (out Uint8x16) {
	saturatingAddU8x16(&out, &a, &b)
	return out
}

func SaturatingAdd8x16(a, b Int8x16) (out Int8x16) {
	saturatingAdd8x16(&out, &a, &b)
	return out
}

func subU8x16(out, a, b *Uint8x16)
func sub8x16(out, a, b *Int8x16)

func SubU8x16(a, b Uint8x16) (out Uint8x16) {
	subU8x16(&out, &a, &b)
	return out
}

func Sub8x16(a, b Int8x16) (out Int8x16) {
	sub8x16(&out, &a, &b)
	return out
}

func saturatingSubU8x16(out, a, b *Uint8x16)
func saturatingSub8x16(out, a, b *Int8x16)

func SaturatingSubU8x16(a, b Uint8x16) (out Uint8x16) {
	saturatingSubU8x16(&out, &a, &b)
	return out
}

func SaturatingSub8x16(a, b Int8x16) (out Int8x16) {
	saturatingSub8x16(&out, &a, &b)
	return out
}

func andU8x16(out, a, b *Uint8x16)
func and8x16(out, a, b *Int8x16)

func AndU8x16(a, b Uint8x16) (out Uint8x16) {
	andU8x16(&out, &a, &b)
	return out
}

func And8x16(a, b Int8x16) (out Int8x16) {
	and8x16(&out, &a, &b)
	return out
}

func andU16x8(out, a, b *Uint16x8)
func and16x8(out, a, b *Int16x8)

func AndU16x8(a, b Uint16x8) (out Uint16x8) {
	andU16x8(&out, &a, &b)
	return out
}

func And16x8(a, b Int16x8) (out Int16x8) {
	and16x8(&out, &a, &b)
	return out
}

func orU8x16(out, a, b *Uint8x16)
func or8x16(out, a, b *Int8x16)

func OrU8x16(a, b Uint8x16) (out Uint8x16) {
	orU8x16(&out, &a, &b)
	return out
}

func Or8x16(a, b Int8x16) (out Int8x16) {
	or8x16(&out, &a, &b)
	return out
}

func xorU8x16(out, a, b *Uint8x16)
func xor8x16(out, a, b *Int8x16)

func XorU8x16(a, b Uint8x16) (out Uint8x16) {
	xorU8x16(&out, &a, &b)
	return out
}

func Xor8x16(a, b Int8x16) (out Int8x16) {
	xor8x16(&out, &a, &b)
	return out
}

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

func shiftRight16x8(out, a *Uint16x8, idx uint)
func shiftRightU16x8(out, a *Uint16x8, idx uint)

func ShiftRight16x8(a Uint16x8, idx uint) (out Uint16x8) {
	// if idx < 1 || idx > 8 {
	// 	return out
	// }
	shiftRight16x8(&out, &a, idx)
	return out
}

func ShiftRightU16x8(a Uint16x8, idx uint) (out Uint16x8) {
	// if idx < 1 || idx > 16 {
	// 	return out
	// }
	shiftRightU16x8(&out, &a, idx)
	return out
}

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
