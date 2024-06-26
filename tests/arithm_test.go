//go:build sse2 || sse3 || sse4.1 || neon

package tests

import (
	"math"
	"simd"
	"slices"
	"testing"
)

func TestAdd8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
		b [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b: [16]int8{math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8},
		},
		{
			a: [16]int8{math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8},
			b: [16]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceAdd8x16(test.a, test.b)
		got := simd.Add8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestAddU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
		b [16]uint8
	}{
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b: [16]uint8{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		},
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceAddU8x16(test.a, test.b)
		got := simd.AddU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestSaturatingAdd8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
		b [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b: [16]int8{math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8},
		},
		{
			a: [16]int8{math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8},
			b: [16]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceSaturatingAdd8x16(test.a, test.b)
		got := simd.SaturatingAdd8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestSaturatingAddU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
		b [16]uint8
	}{
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b: [16]uint8{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		},
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceSaturatingAddU8x16(test.a, test.b)
		got := simd.SaturatingAddU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestSub8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
		b [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b: [16]int8{math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8},
		},
		{
			a: [16]int8{math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8},
			b: [16]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceSub8x16(test.a, test.b)
		got := simd.Sub8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestSubU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
		b [16]uint8
	}{
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint8{},
			b: [16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceSubU8x16(test.a, test.b)
		got := simd.SubU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestSaturatingSub8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
		b [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8, math.MinInt8},
			b: [16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			a: [16]int8{math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8},
			b: [16]int8{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
			b: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceSaturatingSub8x16(test.a, test.b)
		got := simd.SaturatingSub8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestSaturatingSubU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
		b [16]uint8
	}{
		{
			a: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			b: [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b: [16]uint8{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		},
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
			b: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceSaturatingSubU8x16(test.a, test.b)
		got := simd.SaturatingSubU8x16(test.a, test.b)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Logf("b = %v\n", test.b)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}
