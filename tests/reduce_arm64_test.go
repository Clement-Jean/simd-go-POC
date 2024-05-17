package tests

import (
	"simd"
	"testing"
)

func TestReduceMax8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		got := simd.ReduceMax8x16(test.a)
		expected := referenceReduceMax8x16(test.a)

		if expected != got {
			t.Logf("a = %v\n", test.a)
			t.Fatalf("%v: expected %d, got %d\n", test.a, expected, got)
		}
	}
}

func TestReduceMin8x16(t *testing.T) {
	tests := []struct {
		a [16]int8
	}{
		{
			a: [16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			a: [16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			a: *(*[16]int8)(genInt8Arr(16)),
		},
	}

	for _, test := range tests {
		got := simd.ReduceMin8x16(test.a)
		expected := referenceReduceMin8x16(test.a)

		if expected != got {
			t.Logf("a = %v\n", test.a)
			t.Fatalf("%v: expected %d, got %d\n", test.a, expected, got)
		}
	}
}
