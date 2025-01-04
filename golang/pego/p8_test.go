package pego

import (
	"testing"
)

func TestProblem8Spec(t *testing.T) {
	expected := 5832
	actual := findLargestProduct(PROBLEM_8_INPUT, 4)

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}

func TestProblem8Solution(t *testing.T) {
	expected := 23514624000
	actual := SolveProblem8()

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}
