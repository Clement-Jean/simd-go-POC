package main_test

import (
	"testing"

	"simd"
)

var d *[16]uint8

func BenchmarkAddU8x16(b *testing.B) {
	a := &[16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	c := &[16]uint8{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d = simd.AddU8x16(a, c)
	}
}

func BenchmarkAddU8x16For(b *testing.B) {
	a := [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	c := [16]uint8{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(a); i++ {
			d[i] = a[i] + c[i]
		}
	}
}
