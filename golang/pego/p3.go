package pego

import "com.jfisherdev/pego/utils"

func SolveProblem3() int {
	n := 600851475143
	return findLargestPrime(n)
}

func findLargestPrime(n int) int {
	primeFactors, _ := utils.PrimeFactorization(n)
	return primeFactors[len(primeFactors)-1]
}
