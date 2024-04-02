package simd

import (
	"testing"
)

func TestMax8x16(t *testing.T) {
	a := &[16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b := &[16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	out := &[16]int8{}

	Max8x16(out, a, b)

	// check that `a` didn't change
	for i, d := range a {
		if d != int8(i+1) {
			t.Fatalf("expected %d, got %d", i+1, d)
		}
	}

	// check that `b` didn't change
	for i, d := range b {
		if d != int8(16-i) {
			t.Fatalf("expected %d, got %d", 16-i, d)
		}
	}

	for i, d := range out {
		if d != max(a[i], b[i]) {
			t.Fatalf("expected %d, got %d", max(a[i], b[i]), d)
		}
	}
}

func TestMin8x16(t *testing.T) {
	a := &[16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b := &[16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	out := &[16]int8{}

	Min8x16(out, a, b)

	// check that `a` didn't change
	for i, d := range a {
		if d != int8(i+1) {
			t.Fatalf("expected %d, got %d", i+1, d)
		}
	}

	// check that `b` didn't change
	for i, d := range b {
		if d != int8(16-i) {
			t.Fatalf("expected %d, got %d", 16-i, d)
		}
	}

	for i, d := range out {
		if d != min(a[i], b[i]) {
			t.Fatalf("expected %d, got %d", min(a[i], b[i]), d)
		}
	}
}
