package tests

import (
	"simd"
	"testing"
)

func TestMax8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		got := simd.Max8x16(&test.a)
		expected := referenceMax8x16(&test.a)

		if expected != got {
			t.Fatalf("expected %d, got %d\n", expected, got)
		}
	}
}
