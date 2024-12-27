package utils

import (
	"slices"
	"testing"
)

func TestPrimeFactorizationWhenPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37, 41, 43, 47, 53}
	for _, value := range primes {
		result, _ := PrimeFactorization(value)
		if len(result) > 0 {
			t.Fatalf("%d is prime and should have no prime factors", value)
		}
	}
}

func TestPrimeFactorizationWhenNotPrime(t *testing.T) {
	n := 30
	expected := []int{2, 3, 5}
	actual, _ := PrimeFactorization(n)
	if !slices.Equal(expected, actual) {
		t.Fatalf("%d is the expected prime factorization for %d but got %d", expected, n, actual)
	}
}

func TestPrimeFactorizationGivenIllegalValue(t *testing.T) {
	illegalValues := []int{0, -1}

	for _, value := range illegalValues {
		_, err := PrimeFactorization(value)
		if err == nil {
			t.Fatalf("%d is not an allowed value", value)
		}
	}
}

func TestIsPrimeWhenPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37, 41, 43, 47, 53}

	for _, value := range primes {
		isPrime, _ := IsPrime(value)
		if !isPrime {
			t.Fatalf("%d is prime!", value)
		}
	}
}

func TestIsPrimeWhenNotPrime(t *testing.T) {
	notPrimes := []int{1, 4, 51}

	for _, value := range notPrimes {
		isPrime, _ := IsPrime(value)
		if isPrime {
			t.Fatalf("%d is not prime", value)
		}
	}
}

func TestIsPrimeGivenIllegalValue(t *testing.T) {
	illegalValues := []int{0, -1}

	for _, value := range illegalValues {
		_, err := IsPrime(value)
		if err == nil {
			t.Fatalf("%d is not an allowed value", value)
		}
	}
}

func TestPartition(t *testing.T) {
	expected := []int{}

	actual := Parition(51)

	if !slices.Equal(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
