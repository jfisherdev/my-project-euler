package pego

import pego "com.jfisherdev/pego/utils"

func SolveProblem2() int {
	var sum int = 0

	fs := pego.FibonacciSequence{}

	for {
		term := fs.Next()

		if term > 4_000_000 {
			break
		}

		if term%2 == 0 {
			sum += term
		}
	}

	return sum
}
