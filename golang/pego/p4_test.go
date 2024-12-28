package pego

import (
	"testing"
)

func TestProblem4Spec(t *testing.T) {
	expected := 9009
	actual := findLargestPalindrome(10, 99)

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}

func TestProblem4Solution(t *testing.T) {
	expected := 906609

	actual := SolveProblem4()

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}
