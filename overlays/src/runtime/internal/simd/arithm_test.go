package simd

import "testing"

func TestAdd8x16(t *testing.T) {
	a := &[16]uint8{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	b := &[16]uint8{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	out := &[16]uint8{}

	Add8x16(out, a, b)

	// check that `a` didn't change
	for _, d := range a {
		if d != 2 {
			t.Fatalf("expected 2, got %d", d)
		}
	}

	// check that `b` didn't change
	for _, d := range b {
		if d != 255 {
			t.Fatalf("expected 255, got %d", d)
		}
	}

	for _, d := range out {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}
