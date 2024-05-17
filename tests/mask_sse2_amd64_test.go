//go:build sse2 || sse3 || sse4.1

package tests

import (
	"simd"
	"testing"
)

func TestMoveMaskByte8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceMoveMaskByte8x16(test.a)
		got := simd.MovMaskByte8x16(test.a)

		if expected != got {
			t.Fatalf("expected %d, got %d\n", expected, got)
		}
	}
}

func TestMoveMaskByteU8x16(t *testing.T) {
	tests := []struct {
		a [16]uint8
	}{
		{
			a: *(*[16]uint8)(genUint8Arr(16)),
		},
	}

	for _, test := range tests {
		expected := referenceMoveMaskByteU8x16(test.a)
		got := simd.MovMaskByteU8x16(test.a)

		if expected != got {
			t.Fatalf("expected %d, got %d\n", expected, got)
		}
	}
}
