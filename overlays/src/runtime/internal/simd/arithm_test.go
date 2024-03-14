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

func TestSaturatingAddU8x16(t *testing.T) {
	a := &[16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b := &[16]uint8{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	out := &[16]uint8{}

	SaturatingAddU8x16(out, a, b)

	// check that `a` didn't change
	for _, d := range a {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}

	// check that `b` didn't change
	for _, d := range b {
		if d != 255 {
			t.Fatalf("expected 255, got %d", d)
		}
	}

	for _, d := range out {
		if d != 255 {
			t.Fatalf("expected 255, got %d", d)
		}
	}
}

func TestSaturatingAdd8x16(t *testing.T) {
	a := &[16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b := &[16]int8{127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127}
	out := &[16]int8{}

	SaturatingAdd8x16(out, a, b)

	// check that `a` didn't change
	for _, d := range a {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}

	// check that `b` didn't change
	for _, d := range b {
		if d != 127 {
			t.Fatalf("expected 127, got %d", d)
		}
	}

	for _, d := range out {
		if d != 127 {
			t.Fatalf("expected 127, got %d", d)
		}
	}
}

func TestSub8x16(t *testing.T) {
	a := &[16]uint8{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	b := &[16]uint8{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	out := &[16]uint8{}

	Sub8x16(out, a, b)

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
		if d != 3 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}

func TestSaturatingSubU8x16(t *testing.T) {
	a := &[16]uint8{}
	b := &[16]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	out := &[16]uint8{}

	SaturatingSubU8x16(out, a, b)

	// check that `a` didn't change
	for _, d := range a {
		if d != 0 {
			t.Fatalf("expected 0, got %d", d)
		}
	}

	// check that `b` didn't change
	for _, d := range b {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}

	for _, d := range out {
		if d != 0 {
			t.Fatalf("expected 0, got %d", d)
		}
	}
}

func TestSaturatingSub8x16(t *testing.T) {
	a := &[16]int8{-128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128, -128}
	b := &[16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	out := &[16]int8{}

	SaturatingSub8x16(out, a, b)

	// check that `a` didn't change
	for _, d := range a {
		if d != -128 {
			t.Fatalf("expected -128, got %d", d)
		}
	}

	// check that `b` didn't change
	for _, d := range b {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}

	for _, d := range out {
		if d != -128 {
			t.Fatalf("expected -128, got %d", d)
		}
	}
}
