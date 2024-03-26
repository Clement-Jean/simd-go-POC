package simd

import (
	"testing"
)

func TestMax8x16(t *testing.T) {
	a := &[16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	out := Max8x16(a)

	// check that `a` didn't change
	for i, d := range a {
		if d != uint8(i+1) {
			t.Fatalf("expected %d, got %d", i+1, d)
		}
	}

	if out != 16 {
		t.Fatalf("expected 16, got %d", out)
	}
}

func TestMin8x16(t *testing.T) {
	a := &[16]uint8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	out := Min8x16(a)

	// check that `a` didn't change
	for i, d := range a {
		if d != uint8(16-i) {
			t.Fatalf("expected %d, got %d", 16-i, d)
		}
	}

	if out != 1 {
		t.Fatalf("expected 1, got %d", out)
	}
}
