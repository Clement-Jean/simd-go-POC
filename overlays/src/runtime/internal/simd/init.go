package simd

func Splat8x16(out *[16]int8, n uint8)
func SplatU8x16(out *[16]uint8, n int8)

func Extract8x16(out, a, b *[16]int8, idx uint)
func ExtractU8x16(out, a, b *[16]uint8, idx uint)

func Lookup8x16(out, a *[16]int8, b *[16]uint8)
func LookupU8x16(out, a, b *[16]uint8)
