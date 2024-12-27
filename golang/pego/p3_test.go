package pego

import (
	"testing"
)

func TestProblem3Spec(t *testing.T) {
	n := 13195
	expected := 29
	actual := findLargestPrime(n)

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}

func TestProblem3Solution(t *testing.T) {
	expected := 6857

	actual := SolveProblem3()

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}
