//go:build neon

package tests

import (
	"simd"
	"slices"
	"testing"
)

func TestMax8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
		b [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceMax8x16(test.a, test.b)
		got := simd.Max8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}
