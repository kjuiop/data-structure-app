package main

import "fmt"

func main() {
	numList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := 2
	fmt.Println(solution(numList, n))
}

func solution(num_list []int, n int) [][]int {

	result := make([][]int, 0)
	list := make([]int, n)

	index := 0
	for i := 0; i < len(num_list); i++ {

		list[index] = num_list[i]
		index++

		if (i+1)%n == 0 {
			result = append(result, list)
			list = make([]int, n)
			index = 0
		}
	}

	return result
}
