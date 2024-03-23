package main

import (
	"fmt"
	"strings"
)

// learn/courses/30/lessons/258712
func main() {
	friends := []string{"muzi", "ryan", "frodo", "neo"}
	gifts := []string{"muzi frodo", "muzi frodo", "ryan muzi", "ryan muzi", "ryan muzi", "frodo muzi", "frodo ryan", "neo muzi"}

	fmt.Println(solution(friends, gifts))
}

// 선물을 주고 받은 내역이 있다면 덜 받은 쪽이 다음달에 받는다.
// 선물 주고받은 기록이 없거나 주고 받은 수가 같다면 선물 지수로 비교한다.

// 선물 지수 : 이번 딸까지 친구들에게 준 선물 수 - 받은 선물 수
// 선물 지수가 더 큰 사람이 선물 지수가 더 작은 사람에게 선물을 하나 받는다. (받기만 한 애가 많이 준애한테 준다.)
// 선물 지수가 같다면 선물을 주고받지 않는다.
func solution(friends []string, gifts []string) int {

	giftRecord := make(map[string]map[string]int)
	giftScore := make(map[string]int)
	nextGiftRecord := make(map[string]int)
	for _, name := range friends {
		giftRecord[name] = make(map[string]int)
		giftScore[name] = 0
		nextGiftRecord[name] = 0
	}

	for _, str := range gifts {
		arrays := strings.Split(str, " ")
		giver := arrays[0]
		receiver := arrays[1]

		giftRecord[giver][receiver]++
		giftScore[giver]++
		giftScore[receiver]--
	}

	for _, giver := range friends {
		for _, receiver := range friends {
			if giver == receiver {
				continue
			}

			giftsFromGiver := giftRecord[giver][receiver]
			giftsFromReceiver := giftRecord[receiver][giver]

			if giftsFromGiver > giftsFromReceiver {
				nextGiftRecord[giver]++
			} else if giftsFromGiver == giftsFromReceiver && giftScore[giver] > giftScore[receiver] {
				nextGiftRecord[giver]++
			}
		}
	}

	max := 0
	for _, cnt := range nextGiftRecord {
		if max <= cnt {
			max = cnt
		}
	}

	return max
}
