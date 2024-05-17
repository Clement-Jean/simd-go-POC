//go:build sse2 || sse3 || sse4.1

package tests

import (
	"simd"
	"slices"
	"testing"
)

func TestAnd16x8(t *testing.T) {
	tests := []struct {
		a [8]int16
		b [8]int16
	}{
		{
			a: [8]int16{-1, -2, -3, -4, -5, -6, -7, -8},
			b: [8]int16{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			a: [8]int16{-1, -2, -3, -4, -5, -6, -7, -8},
			b: [8]int16{16, 15, 14, 13, 12, 11, 10, 9},
		},
	}

	for _, test := range tests {
		expected := referenceAnd16x8(test.a, test.b)
		got := simd.And16x8(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestAndU16x8(t *testing.T) {
	tests := []struct {
		a [8]uint16
		b [8]uint16
	}{
		{
			a: [8]uint16{1, 2, 3, 4, 5, 6, 7, 8},
			b: [8]uint16{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			a: [8]uint16{1, 2, 3, 4, 5, 6, 7, 8},
			b: [8]uint16{16, 15, 14, 13, 12, 11, 10, 9},
		},
	}

	for _, test := range tests {
		expected := referenceAndU16x8(test.a, test.b)
		got := simd.AndU16x8(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}
