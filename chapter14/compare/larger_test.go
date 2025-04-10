package compare

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	a, b := 2, 1
	want := 2
	got := Larger(a,b)
	if got != want {
		t.Error(errorString(a, b, want, got))
	}
}

func TestSecondLarger(t *testing.T) {
	a, b := 4, 8
	want := 8
	got := Larger(4,8)
	if got != want {
		t.Error(errorString(a, b, want, got))
	}
}

func errorString(a, b, want, got int) string {
	return fmt.Sprintf("Larger(%d, %d) = \"%d\", want \"%d\"",a , b, want, got)
}
