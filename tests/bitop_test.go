//go:build sse2 || sse3 || sse4.1 || neon

package tests

import (
	"simd"
	"slices"
	"testing"
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
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceAnd8x16(test.a, test.b)
		got := simd.And8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
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
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceAndU8x16(test.a, test.b)
		got := simd.AndU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
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
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceOr8x16(test.a, test.b)
		got := simd.Or8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
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
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceOrU8x16(test.a, test.b)
		got := simd.OrU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
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
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceXor8x16(test.a, test.b)
		got := simd.Xor8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
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
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceXorU8x16(test.a, test.b)
		got := simd.XorU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}
