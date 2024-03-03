package main

import "fmt"

func main() {
	fmt.Println("expected : [2, 1], actual : ", solution([]int{93, 30, 55}, []int{1, 30, 5}))
	fmt.Println("expected : [1, 3, 2], actual : ", solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1}))
}

func solution(progresses []int, speeds []int) []int {
	if len(progresses) < 0 || len(progresses) > 100 || len(speeds) < 0 || len(speeds) > 100 {
		return nil
	}

	if len(progresses) != len(speeds) {
		return nil
	}

	var max, count int = 0, 1
	var result []int

	for idx, progress := range progresses {
		var remainPeriod int = 100 - progress
		var a int = remainPeriod / speeds[idx]

		if max > a {
			result[len(result)-1]++
			continue
		}
		result = append(result, count)
		max = a
	}

	return result
}

func answer(progresses []int, speeds []int) []int {

	var max, count int = 0, 0
	var result []int

	for idx, completeTime := range progresses {
		var remainPeriod int = 100 - completeTime
		var completePeriod int = remainPeriod / speeds[idx]

		if remainPeriod%speeds[idx] != 0 {
			completePeriod++
		}

		if max < completePeriod {
			result = append(result, count)
			count = 1
			max = completePeriod
		} else {
			count++
		}
	}

	result = append(result[:0], result[1:]...)
	return append(result, count)
}

func problem(progresses []int, speeds []int) []int {
	if len(progresses) < 0 || len(progresses) > 100 || len(speeds) < 0 || len(speeds) > 100 {
		return nil
	}

	if len(progresses) != len(speeds) {
		return nil
	}

	result := make(map[int]int)
	compare := 0
	for i := 0; i < len(progresses); i++ {
		remain := 100 - progresses[i]
		a := remain / speeds[i]

		if len(result) == 0 {
			result[0]++
			compare = a
			continue
		}

		resultIndex := len(result)
		if compare > a {
			result[resultIndex-1]++
			continue
		}
		result[resultIndex]++
		compare = a
	}

	answer := make([]int, 0)
	for _, val := range result {
		answer = append(answer, val)
	}

	return answer
}
