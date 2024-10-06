package main

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/142086
func main() {
	fmt.Println(solution("banana"))
	fmt.Println(solution("foobar"))
}

// 1. 바로 앞의 가장 가까운 알파벳의 위치를 구한다.
// -> break, 중첩 for 문, 결과값 저장
func solution(s string) []int {

	var result []int
	arrays := strings.Split(s, "")
	for i := 0; i < len(arrays); i++ {
		if i == 0 {
			result = append(result, -1)
			continue
		}
		index := 0
		for j := 0; j < i; j++ {
			if arrays[i] == arrays[j] {
				index = i - j
			}
		}

		if index == 0 {
			result = append(result, -1)
		} else {
			result = append(result, index)
		}
	}

	return result
}

func solution1(s string) []int {
	m := make(map[string]int)
	var res []int

	for i := range s {
		val, ok := m[string(s[i])]
		if !ok {
			m[string(s[i])] = i
			res = append(res, -1)
		} else {
			m[string(s[i])] = i
			res = append(res, i-val)
		}
	}
	return res
}

func solution2(s string) []int {
	r := make([]int, 0)

	for i := range s {
		index := strings.LastIndex(s[:i], s[i:i+1])
		if index != -1 {
			index = i - index
		}

		r = append(r, index)
	}

	return r
}
