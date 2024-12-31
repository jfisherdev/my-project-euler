package pego

import (
	"testing"
)

func TestProblem5Spec(t *testing.T) {
	expected := 2520
	actual := smallestMultiple(10)

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}

func TestProblem5Solution(t *testing.T) {
	expected := 232792560

	actual := SolveProblem5()

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}
