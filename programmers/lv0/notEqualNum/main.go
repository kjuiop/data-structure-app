package main

import "fmt"

func main() {
	fmt.Println("expected : [1,3,0,1], actual : ", solution([]int{1, 1, 3, 3, 0, 1, 1}))
	fmt.Println("expected : [4,3], actual : ", solution([]int{4, 4, 4, 3, 3}))
}

func solution(nums []int) []int {

	if len(nums) < 0 || len(nums) > 1000000 {
		return nil
	}

	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		result = append(result, nums[i])
	}

	return result
}
