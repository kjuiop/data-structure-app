package main

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/155652
func main() {
	fmt.Println(solution("aukks", "wbqd", 5))
}

func solution(s, skip string, index int) string {
	var result []rune
	for j, char := range s {
		for i := 0; i < index; {
			char++
			if char > 'z' { // 알파벳의 범위를 넘어가면 다시 'a'로 돌아감
				char = 'a'
			}
			// skip 문자열에 포함되지 않는 문자만 count
			if !strings.ContainsRune(skip, char) {
				i++
			}
			fmt.Printf("Sub Routine Character at index %d: %c (ASCII: %d)\n", j, char, char)
		}
		fmt.Printf("Final Character at index %d: %c (ASCII: %d)\n", j, char, char)
		result = append(result, char)
	}

	return string(result)
}
