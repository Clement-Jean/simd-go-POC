package simd

func Add8x16(out, a, b *[16]uint8)

func SaturatingAddU8x16(out, a, b *[16]uint8)
func SaturatingAdd8x16(out, a, b *[16]int8)

func Sub8x16(out, a, b *[16]uint8)

func SaturatingSubU8x16(out, a, b *[16]uint8)
func SaturatingSub8x16(out, a, b *[16]int8)

func And8x16(out, a, b *[16]uint8)
