//go:build sse2 || sse3 || sse4.1 || neon

package tests

import (
	"simd"
	"slices"
	"testing"
)

func TestMaxU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
		b [16]uint8
	}{
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceMaxU8x16(test.a, test.b)
		got := simd.MaxU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestMinU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
		b [16]uint8
	}{
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceMinU8x16(test.a, test.b)
		got := simd.MinU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}
