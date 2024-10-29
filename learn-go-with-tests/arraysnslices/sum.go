package main

func Sum(n []int) int {
	sum := 0

	for _, v := range n {
		sum += v
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = Sum(n)
	}

	return result
}

func SumAllTails(numbers ...[]int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		if len(n) == 0 {
			result[i] = 0
		} else {
			result[i] = Sum(n[1:])
		}
	}

	return result
}
