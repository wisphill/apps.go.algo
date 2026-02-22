package main

import "fmt"

func two_sum(arr []int, total int) [][]int {
	// dictionay
	dict := make(map[int]int)
	res := make([][]int, 0)
	for _, value := range arr {
		dict[value] = dict[value] + 1
	}

	for _, value := range arr {
		remaining_value := total - value
		if dict[remaining_value] > 0 && dict[value] > 0 {
			res = append(res, []int{value, remaining_value})
			println("values: ", value, dict[value])
		}
	}

	return res
}

func main() {
	couple := two_sum([]int{1, 2, 3, 8, 13, 3, 3, 2, 1}, 5)
	for _, values := range couple {
		fmt.Println(values[0], values[1])
	}
}
