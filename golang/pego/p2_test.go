package pego

import (
	"testing"
)

func TestProblem2Solution(t *testing.T) {
	expected := 4613732

	actual := SolveProblem2()

	if expected != actual {
		t.Fatalf("Should get %d got %d", expected, actual)
	}
}
