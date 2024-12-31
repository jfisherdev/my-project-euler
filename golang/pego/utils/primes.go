package utils

import (
	"errors"
	"math"
)

func PrimeFactorization(n int) ([]int, error) {
	var primeFactors []int

	if n <= 0 {
		return nil, errors.New("value must be >= 1")
	} else if n > 2 {
		var isPrime, _ = IsPrime(n)
		if !isPrime {
			midpoint := int(math.Floor(math.Sqrt(float64(n))))
			for candidate := 2; candidate <= midpoint; candidate++ {
				if n%candidate == 0 {
					isPrimeFactor, _ := IsPrime(candidate)
					if isPrimeFactor {
						primeFactors = append(primeFactors, candidate)
					}
				}
			}
		}
	}

	return primeFactors, nil
}

func IsPrime(n int) (bool, error) {

	if n <= 0 {
		return false, errors.New("value must be >= 1")
	}

	if n == 1 {
		return false, nil
	}

	if n == 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	midpoint := int(math.Floor(math.Sqrt(float64(n))))

	for candidate := 2; candidate <= midpoint; candidate++ {
		if n%candidate == 0 {
			return false, nil
		}
	}

	return true, nil
}

func Parition(n int) []int {
	var partitions []int
	var midpoint = n / 2

	for {
		if midpoint < 2 {
			break
		}
		partitions = append(partitions, midpoint)
		midpoint = midpoint / 2
	}

	return partitions
}
