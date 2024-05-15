package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/159994?language=go
func main() {
	fmt.Println(solution([]string{"i", "drink", "water"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"}))
	fmt.Println(solution([]string{"i", "water", "drink"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"}))
	fmt.Println(solution([]string{"a", "b", "c"}, []string{"d", "e", "f"}, []string{"a", "d", "f"}))
}

func solution(cards1 []string, cards2 []string, goal []string) string {
	// 원하는 카드 뭉치에서 카드를 순서대로 한 장씩 사용
	// 한 번 사용한 카드는 다시 사용 못함
	// 카드를 사용하지 않고는 다음 카드로 넘어갈 수 없음
	// 기존에 주어진 카드 뭉치의 순서는 바꿀 수 없음

	result := "Yes"
	i, j := 0, 0
	for _, s := range goal {

		if i < len(cards1) && s == cards1[i] {
			i++
			continue
		}

		if j < len(cards2) && s == cards2[j] {
			j++
			continue
		}

		result = "No"
		break
	}

	return result
}
