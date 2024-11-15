package main

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/134240
func main() {
	fmt.Println(solution([]int{1, 3, 4, 6}))
}

func solution(food []int) string {
	// 선수들은 1대1 대결을 한다.
	// 한명은 왼쪽부터, 한명은 오른쪽부터 음식을 순서대로 먹는다.
	// 칼로리가 낮은 음식을 먼저 먹을 수 있게 배치
	// 대회에서 사용하지 못하는 음식이 생길 수 있음
	// result 는 음식 배열

	// 0을 기준으로 절반만 구하고, 나머지는 대칭으로 뒤집는다.
	var leftSide strings.Builder

	// 각 음식을 절반씩 나눠 왼쪽 선수의 배치 생성
	for i := 1; i < len(food); i++ {
		half := food[i] / 2
		leftSide.WriteString(strings.Repeat(fmt.Sprintf("%d", i), half))
	}

	// 왼쪽 배치 가져오기
	left := leftSide.String()

	// 결과: 왼쪽 + 물(0) + 오른쪽(왼쪽을 뒤집은 형태)
	result := left + "0" + reverseString(left)

	return result
}

// 문자열 뒤집기 함수
func reverseString(s string) string {
	n := len(s)
	reversed := make([]byte, n)
	for i := 0; i < n; i++ {
		reversed[i] = s[n-1-i]
	}
	return string(reversed)
}
