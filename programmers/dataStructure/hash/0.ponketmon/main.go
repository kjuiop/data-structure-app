package main

import "fmt"

func main() {
	fmt.Println("expected : 2, actual : ", solution([]int{3, 1, 2, 3}))
	fmt.Println("expected : 3, actual : ", solution([]int{3, 3, 3, 2, 2, 4}))
	fmt.Println("expected : 2, actual : ", solution([]int{3, 3, 3, 2, 2, 2}))
}

func solution(nums []int) int {

	totalCnt := len(nums)
	if totalCnt < 1 || totalCnt > 10000 {
		return -1
	}

	if totalCnt%2 != 0 {
		return -1
	}

	pocketBox := make(map[int]bool)

	for _, num := range nums {

		if num < 1 || num > 200000 {
			return -1
		}

		if len(pocketBox) >= totalCnt/2 {
			break
		}

		pocketBox[num] = true
	}

	return len(pocketBox)
}
