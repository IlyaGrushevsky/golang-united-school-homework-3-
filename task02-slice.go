package homework

import "sort"

func reverse(input []int64) (result []int64) {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})
	result = input
	return result
}
