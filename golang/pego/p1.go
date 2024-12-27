package pego

import (
	"slices"
)

func SolveProblem1() int {
	limit := 1000
	values := []int{3, 5}
	return GetMultiplesForValues(limit, values)
}

func GetMultiplesForValues(limit int, values []int) int {
	sum := 0
	var multiples []int
	for _, value := range values {
		var multiplesForValue []int = getMultiplesForValue(limit, value)
		for _, multipleForValue := range multiplesForValue {
			if !slices.Contains(multiples, multipleForValue) {
				multiples = append(multiples, multipleForValue)
			}
		}
		//sum += getMultiplesForValue(limit, value)
	}

	for _, multiple := range multiples {
		sum += multiple
	}

	return sum
}

func getMultiplesForValue(limit int, value int) []int {
	var multiples []int
	multiplier := 1
	for {
		multiple := multiplier * value
		if multiple >= limit {
			break
		}
		multiples = append(multiples, multiple)
		multiplier++
	}
	return multiples
}
