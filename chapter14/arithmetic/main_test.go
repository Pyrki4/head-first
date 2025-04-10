package arithmetic

import "testing"

func TestAdd (t *testing.T) {
	if Add(1, 2) != 3.0 {
		t.Error("1 + 2 did not equals 3")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(8,4) != 4 {
		t.Error("8 - 4 did not equals 4")
	}
}
