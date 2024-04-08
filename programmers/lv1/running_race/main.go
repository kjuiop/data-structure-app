package main

import "fmt"

// learn/courses/30/lessons/178871
func main() {
	fmt.Println(solution([]string{"mumu", "soe", "poe", "kai", "mine"}, []string{"kai", "kai", "mine", "mine"}))
}

// rank 결과를 맵으로 들고 있는 다음 바로 접근해서 찾기
func solution(players []string, callings []string) []string {

	rank := make(map[string]int, 0)
	for idx, player := range players {
		rank[player] = idx
	}

	for _, calling := range callings {
		index := rank[calling]

		temp := players[index-1]

		players[index-1] = players[index]
		players[index] = temp

		rank[players[index-1]]--
		rank[players[index]]++
	}

	return players
}

/**
단순 중첩 for 문
func solution(players []string, callings []string) []string {

	for _, calling := range callings {
		for j, player := range players {
			var temp string
			if calling == player {
				temp = players[j-1]
				players[j-1] = player
				players[j] = temp
				break
			}
		}
	}

	return players
}
*/
