package tests

import (
	"testing"

	"simd"
)

func TestAnd8x16(t *testing.T) {
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
	}

	for _, test := range tests {
		expected := referenceAnd8x16(&test.a, &test.b)
		got := simd.And8x16(&test.a, &test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %d at index %d, got %d\n", expected[i], i, got[i])
			}
		}
	}
}

func TestAndU8x16(t *testing.T) {
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
	}

	for _, test := range tests {
		expected := referenceAndU8x16(&test.a, &test.b)
		got := simd.AndU8x16(&test.a, &test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %d at index %d, got %d\n", expected[i], i, got[i])
			}
		}
	}
}

func TestOr8x16(t *testing.T) {
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
	}

	for _, test := range tests {
		expected := referenceOr8x16(&test.a, &test.b)
		got := simd.Or8x16(&test.a, &test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %d at index %d, got %d\n", expected[i], i, got[i])
			}
		}
	}
}

func TestOrU8x16(t *testing.T) {
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
	}

	for _, test := range tests {
		expected := referenceOrU8x16(&test.a, &test.b)
		got := simd.OrU8x16(&test.a, &test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %d at index %d, got %d\n", expected[i], i, got[i])
			}
		}
	}
}

func TestXor8x16(t *testing.T) {
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
	}

	for _, test := range tests {
		expected := referenceXor8x16(&test.a, &test.b)
		got := simd.Xor8x16(&test.a, &test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %d at index %d, got %d\n", expected[i], i, got[i])
			}
		}
	}
}

func TestXorU8x16(t *testing.T) {
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
	}

	for _, test := range tests {
		expected := referenceXorU8x16(&test.a, &test.b)
		got := simd.XorU8x16(&test.a, &test.b)

		for i := 0; i < len(got); i++ {
			if expected[i] != got[i] {
				t.Fatalf("expected %d at index %d, got %d\n", expected[i], i, got[i])
			}
		}
	}
}
