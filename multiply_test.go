package multiply

import "testing"

func TestMultiply(t *testing.T) {
	t.Parallel()

	got := Multiply(2, 2)
	want := 4

	if got != want {
		t.Errorf("Multiply(2, 2) = %d; want %d", got, want)
	}
}
