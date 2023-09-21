package main

import "fmt"

func main() {
	fmt.Println("expected : 2, actual : ", solution([]int{3, 1, 2, 3}))
	fmt.Println("expected : 3, actual : ", solution([]int{3, 3, 3, 2, 2, 4}))
	fmt.Println("expected : 2, actual : ", solution([]int{3, 3, 3, 2, 2, 2}))
}

func solution(nums []int) int {

	// 1. n/2 를 가져가도 된다고?
	// 각 종류별로 포켓몬을 다양하게 가져갈 수 있는지에 대한 최대값

	totalCnt := len(nums)
	if totalCnt < 1 || totalCnt > 10000 {
		return -1
	}

	if totalCnt%2 != 0 {
		return -1
	}

	pocketBox := make(map[int]bool)

	for _, val := range nums {
		if len(pocketBox) >= totalCnt/2 {
			break
		}
		pocketBox[val] = true
	}

	return len(pocketBox)
}
