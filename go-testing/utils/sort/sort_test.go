package sort

import (
	"fmt"
	"testing"
)

func TestBUbbleSortNoError(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first emelemt should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
	fmt.Println(elements)
}
