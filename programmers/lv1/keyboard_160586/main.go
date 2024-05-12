package main

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/160586
func main() {
	fmt.Println(solution([]string{"ABACD", "BCEFD"}, []string{"ABCD", "AABB"}))
	fmt.Println(solution([]string{"AA"}, []string{"B"}))
	fmt.Println(solution([]string{"AGZ", "BSSS"}, []string{"ASA", "BGZ"}))
}

func solution(keymap []string, targets []string) []int {

	// 특정 문자열을 작성할 때, 키를 최소 몇 번 눌러야 그 문자열을 작성할 수 있는지 값을 구하기
	// keymap : 자판
	// target : 입력하려고 하는 문자

	a := make(map[string]int)
	for _, key := range keymap {
		str := strings.Split(key, "")
		for idx, s := range str {

			val := idx + 1
			if _, exist := a[s]; !exist {
				a[s] = val
				continue
			}

			if val < a[s] {
				a[s] = val
			}
		}
	}

	var answer []int
	for _, t := range targets {
		result := 0
		str := strings.Split(t, "")
		for _, s := range str {
			if _, exist := a[s]; !exist {
				result = -1
				break
			}
			result += a[s]
		}
		answer = append(answer, result)
	}

	return answer
}
