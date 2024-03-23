package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	var year, month int

	fmt.Print("조사하고자 하는 연도를 입력하세요: ")
	if _, err := fmt.Scan(&year); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	fmt.Print("조사하고자 하는 월을 입력하세요: ")
	if _, err := fmt.Scan(&month); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	monthType := time.Month(month)

	// 해당 연도와 월의 마지막 날짜를 가져옵니다.
	lastDay := time.Date(year, monthType, 0, 0, 0, 0, 0, time.UTC)

	// 결과 출력
	fmt.Printf("%d년 %s월은 %d일까지 있습니다.\n", year, monthType.String(), lastDay.Day())
}
