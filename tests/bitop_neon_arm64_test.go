//go:build neon

package tests

import (
	"simd"
	"slices"
	"testing"
)

func TestShiftRight8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceShiftRight8x16(test.a, 4)
		// this test only support 4 because C intrinsic wouldn't accept a non-const
		got := simd.ShiftRight8x16(test.a, 4)

		if !slices.Equal(expected[:], got[:]) {
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestShiftRightU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
	}{
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceShiftRightU8x16(test.a, 4)
		// this test only support 4 because C intrinsic wouldn't accept a non-const
		got := simd.ShiftRightU8x16(test.a, 4)

		if !slices.Equal(expected[:], got[:]) {
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestShiftLeft8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceShiftLeft8x16(test.a, 4)
		// this test only support 4 because C intrinsic wouldn't accept a non-const
		got := simd.ShiftLeft8x16(test.a, 4)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}

func TestShiftLeftU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
	}{
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceShiftLeftU8x16(test.a, 4)
		// this test only support 4 because C intrinsic wouldn't accept a non-const
		got := simd.ShiftLeftU8x16(test.a, 4)

		if !slices.Equal(expected[:], got[:]) {
			t.Logf("a = %v\n", test.a)
			t.Fatalf("expected %v, got %v\n", expected, got)
		}
	}
}
