package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/172928
func main() {
	fmt.Println(solution(8, 4, []int{2, 3, 6}))
	fmt.Println(solution(5, 4, []int{1, 3}))
	fmt.Println(solution(4, 1, []int{1, 2, 3, 4}))
}

func solution(n int, m int, section []int) int {

	// n 은 칠해야 하는 최종 길이
	// 롤러의 길이 : m
	// 롤러로 페인트칠을 하는 횟수를 최소화하기 위함
	// 정수 n, m 과 다시 페인트를 칠하기로 정한 구역들의 번호가 담긴 정수 배열 section

	tryCnt := 0
	temp := 0
	for idx, val := range section {

		if idx == 0 {
			temp = val + m - 1
			tryCnt++
			continue
		}

		if temp > n {
			break
		}

		if val > temp {
			temp = val + m - 1
			tryCnt++
		}
	}

	return tryCnt
}
