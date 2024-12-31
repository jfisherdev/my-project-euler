package pego

func SolveProblem6() int {
	return sqaureOfSum(100) - sumOfSquares(100)
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

func sqaureOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
