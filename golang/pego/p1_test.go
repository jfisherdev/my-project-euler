package pego

import (
	"testing"
)

func TestProblem1Spec(t *testing.T) {
	limit := 10
	values := []int{3, 5}

	expected := 23

	actual := GetMultiplesForValues(limit, values)

	if expected != actual {
		t.Fatalf("Should get %d given %d and %d but got %d", expected, limit, values, actual)
	}
}

func TestProblem1Solution(t *testing.T) {
	limit := 1000
	values := []int{3, 5}

	expected := 233168

	actual := SolveProblem1()

	if expected != actual {
		t.Fatalf("Should get %d given %d and %d but got %d", expected, limit, values, actual)
	}
}
