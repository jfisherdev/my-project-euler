package pego

import (
	"testing"

	"com.jfisherdev/pego/utils"
)

func TestProblem7Spec(t *testing.T) {
	expected := 13
	actual := utils.NthPrime(6)

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}

func TestProblem7Solution(t *testing.T) {
	expected := 104743
	actual := SolveProblem7()

	if expected != actual {
		t.Fatalf("Should get %d, but got %d", expected, actual)
	}
}
