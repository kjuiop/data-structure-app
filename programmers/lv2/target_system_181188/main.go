package main

import (
	"fmt"
	"sort"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/181188
func main() {
	fmt.Println(solution([][]int{{4, 5}, {4, 8}, {10, 14}, {11, 13}, {5, 12}, {3, 7}, {1, 4}}))
}

func solution(targets [][]int) int {
	cnt := 0

	fmt.Println(targets)
	// 한 발 맞출 때 최대한 많은 미사일을 맞추면 된다.
	// start 와 end 값을 확인하고 끊겨지는 구간을 확인하면 된다.
	// 값을 비교하기 위해서는 정렬이 필요하다.
	sort.Slice(targets[:], func(i, j int) bool {
		return targets[i][1] < targets[j][1]
	})

	fmt.Println(targets)

	// 끝점을 저장
	endPoint := 0
	for _, target := range targets {
		// 시작점이 endpoint 보다 작다면 이전 미사일로 격추할 수 있음
		if target[0] < endPoint {
			continue
		} else {
			cnt++
			endPoint = target[1]
		}
	}

	return cnt
}
