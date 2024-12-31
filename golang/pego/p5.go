package pego

import "com.jfisherdev/pego/utils"

func SolveProblem5() int {
	return smallestMultiple(20)
}

func smallestMultiple(limit int) int {
	var result int = 0
	var primeVals []int = []int{}
	for i := 1; i <= limit; i++ {
		isPrime, _ := utils.IsPrime(i)
		if isPrime {
			primeVals = append(primeVals, i)
		}
	}

	if len(primeVals) > 0 {
		var candidate int = primeVals[len(primeVals)-1] + 1
		for {
			allDivisible := true
			for value := 1; value <= limit; value++ {
				if !(candidate%value == 0) {
					allDivisible = false
					break
				}
			}
			if allDivisible {
				result = candidate
				break
			} else {
				candidate++
			}
		}
	}

	return result
}
