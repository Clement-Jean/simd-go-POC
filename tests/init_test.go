//go:build ssse3 || neon

package tests

import (
	"simd"
	"testing"
)

func TestExtract8x16(t *testing.T) {
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
		expected := referenceExtract8x16(test.a, test.b, 15)
		got := simd.Extract8x16(test.a, test.b, 15)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %v, got %v\n", expected, got)
			}
		}
	}
}

func TestExtractU8x16(t *testing.T) {
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
		expected := referenceExtractU8x16(test.a, test.b, 15)
		got := simd.ExtractU8x16(test.a, test.b, 15)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %v, got %v\n", expected, got)
			}
		}
	}
}

func TestLookup8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
		b [16]uint8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceLookup8x16(test.a, test.b)
		got := simd.Lookup8x16(test.a, test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %v, got %v\n", expected, got)
			}
		}
	}
}

func TestLookupU8x16(t *testing.T) {
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
		expected := referenceLookupU8x16(test.a, test.b)
		got := simd.LookupU8x16(test.a, test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %v, got %v\n", expected, got)
			}
		}
	}
}
