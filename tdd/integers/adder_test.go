package integers

import "testing"

// Add takes two integers and returns the sum of them.
func Add(a int, b int) int {
	return a + b
}

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	want := 6

	if sum != want {
		t.Errorf("expected '%d' but got '%d'", want, sum)
	}
}
