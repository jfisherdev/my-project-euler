package pego

import (
	"testing"
)

func TestProblem6Spec(t *testing.T) {
	expected := 2640
	actual := sqaureOfSum(10) - sumOfSquares(10)

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}

func TestProblem6Solution(t *testing.T) {
	expected := 25164150

	actual := SolveProblem6()

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}
