package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/176963
func main() {
	fmt.Println(solution([]string{"may", "kein", "kain", "radi"}, []int{5, 10, 1, 3}, [][]string{{"may", "kein", "kain", "radi"}, {"may", "kein", "brin", "deny"}, {"kon", "kain", "may", "coni"}}))
}

func solution(name []string, yearning []int, photo [][]string) []int {
	score := make(map[string]int)
	for i, v := range name {
		score[v] = yearning[i]
	}
	var photoScore []int
	for _, array := range photo {
		total := 0
		for _, n := range array {
			total += score[n]
		}
		photoScore = append(photoScore, total)
	}
	return photoScore
}
