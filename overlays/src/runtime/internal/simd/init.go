package simd

func Extract8x16(out, a, b *[16]int8, idx uint)
func ExtractU8x16(out, a, b *[16]uint8, idx uint)

func Lookup8x16(out, a, b *[16]int8)
func LookupU8x16(out, a, b *[16]uint8)
