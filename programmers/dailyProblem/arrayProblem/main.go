package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 5, 2, 6, 3, 7, 4}

	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	fmt.Printf("result : %v\n", solution(array, commands))
}

func solution(array []int, commands [][]int) []int {

	// 2번째에서 5번째까지 자른 후 정렬합니다. 세번째 숫자는 5입니다.
	result := make([]int, 0)

	for _, val := range commands {

		copiedArray := make([]int, len(array))
		copy(copiedArray, array)

		startIndex := val[0] - 1
		endIndex := val[1]
		resultIndex := val[2] - 1

		answer := array[startIndex:endIndex]
		fmt.Printf("array : %d\n", array[3])
		fmt.Printf("startIndex : %d, endIndex : %d, resultIndex : %d, answer : %d\n", startIndex, endIndex, resultIndex, answer)

		sort.Ints(answer)

		result = append(result, answer[resultIndex])
	}

	return result
}
