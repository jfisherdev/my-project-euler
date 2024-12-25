package pego

import (
	"slices"
	"testing"
)

func TestGet20Terms(t *testing.T) {
	limit := 20

	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}
	fs := FibonacciSequence{}
	for n := 0; n < limit; n++ {
		fs.Next()
	}

	actual := fs.Sequence
	if !slices.Equal(expected, actual) {
		t.Fatalf("Should get %d but got %d", expected, actual)
	}
}
