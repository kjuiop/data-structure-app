package main

import "fmt"

func main() {
	fmt.Println("expected : 1, actual : ", solution([][]int{{1, 4}, {9, 2}, {3, 8}, {11, 6}}))
	fmt.Println("expected : 0, actual : ", solution([][]int{{3, 5}, {4, 1}, {2, 4}, {5, 10}}))
}

func solution(dots [][]int) int {
	return 0
}
