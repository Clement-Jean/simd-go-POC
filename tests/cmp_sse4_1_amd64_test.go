//go:build sse4.1

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

func TestMin8x16(t *testing.T) {
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
		expected := referenceMin8x16(test.a, test.b)
		got := simd.Min8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestAllZerosU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
	}{
		{
			a: [16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			a: [16]uint8{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{
			a: [16]uint8{},
		},
	}

	for _, test := range tests {
		expected := referenceAllZerosU8x16(test.a)
		got := simd.AllZerosU8x16(test.a)

		if expected != got {
			t.Logf("a = %v\n", test.a)
			t.Fatalf("expected %t, got %t\n", expected, got)
		}
	}
}

func TestAllZeros8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: [16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			a: [16]int8{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{
			a: [16]int8{},
		},
	}

	for _, test := range tests {
		expected := referenceAllZeros8x16(test.a)
		got := simd.AllZeros8x16(test.a)

		if expected != got {
			t.Logf("a = %v\n", test.a)
			t.Fatalf("expected %t, got %t\n", expected, got)
		}
	}
}
