package main

func Sum(a []int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}
	return sum
}

func SumAll(s ...[]int) []int {
	var result = make([]int, len(s))

	for i, val := range s {
		result[i] = Sum(val)
	}

	return result
}

func SumAllTails(s ...[]int) []int {
	var result []int

	for _, val := range s {
		var newItem int
		if len(val) > 1 {
			newItem = Sum(val[1:])
		} else {
			newItem = 0
		}
		result = append(result, newItem)
	}

	return result
}
