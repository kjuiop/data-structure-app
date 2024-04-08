package main

import (
	"fmt"
	"log"
)

func main() {
	var houseCnt int
	var peopleCnt []int

	fmt.Print("집의 갯수를 입력해주세요.: ")
	if _, err := fmt.Scan(&houseCnt); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	fmt.Print("한 집에 살고 있는 사람의 수를 입력해주세요.: ")
	for i := 0; i < houseCnt; i++ {
		var num int
		if _, err := fmt.Scan(&num); err != nil {
			log.Fatalf("정수를 입력해주세요.")
		}
		peopleCnt = append(peopleCnt, num)
	}

	fmt.Printf("houseCnt %d, peopleCnt %v\n", houseCnt, peopleCnt)

	// 1 2 3 2 6

	fmt.Println(solution(houseCnt, peopleCnt))
}

func solution(houseCnt int, peopleCnt []int) int {
	var sum []int
	for i := 0; i < houseCnt; i++ {
		total := 0
		for j, people := range peopleCnt {
			if i == j {
				continue
			}
			diff := j - i
			if diff < 0 {
				diff *= -1
			}
			total += diff * people
		}
		sum = append(sum, total)
	}

	fmt.Println(sum)

	return findMin(sum)
}

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0 // 빈 슬라이스인 경우 0을 반환하거나, 에러 처리를 수행할 수 있습니다.
	}

	minValue := arr[0] // 초기값으로 첫 번째 요소를 설정

	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}
