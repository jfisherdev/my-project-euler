package pego

import (
	"strconv"
	"strings"
)

func SolveProblem4() int {
	return findLargestPalindrome(100, 999)
}

func findLargestPalindrome(min int, max int) int {
	largestPalindrome := 0
	for i := max; i >= min; i-- {
		for j := max; j >= min; j-- {
			candidate := i * j
			if isPalindrome(candidate) && candidate > largestPalindrome {
				largestPalindrome = candidate
			}
		}
	}
	return largestPalindrome
}

func isPalindrome(n int) bool {
	asString := strconv.Itoa(n)
	reverseBuilder := strings.Builder{}
	for i := len(asString) - 1; i >= 0; i-- {
		reverseBuilder.WriteByte(asString[i])
	}

	reversed := reverseBuilder.String()
	return asString == reversed
}
