package main

import (
	"fmt"
	"strconv"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/147355
func main() {
	fmt.Println(solution("3141592", "271"))
	fmt.Println(solution("500220839878", "7"))
	fmt.Println(solution("10203", "15"))
}

func solution(t string, p string) int {

	pLen := len(p)
	routineCnt := len(t) - pLen
	result := 0
	pNum, _ := strconv.Atoi(p)

	for i := 0; i <= routineCnt; i++ {
		s := t[i : pLen+i]
		sNum, _ := strconv.Atoi(s)
		if sNum <= pNum {
			result++
		}
	}

	return result
}
