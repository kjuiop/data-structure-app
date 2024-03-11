package main

import (
	"fmt"
	"log"
)

func main() {
	var num1, num2 int

	// 사용자로부터 두 수 입력 받기
	fmt.Print("첫 번째 수를 입력하세요: ")
	if _, err := fmt.Scan(&num1); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	fmt.Print("두 번째 수를 입력하세요: ")
	if _, err := fmt.Scan(&num2); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	cnt := 0
	for i := num1; i <= num2; i++ {
		if isDigitSum(i) {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}

func isDigitSum(num int) bool {
	digitSum := (num / 100) + ((num / 10) % 10) + (num % 10)
	if digitSum%2 == 0 {
		return true
	}

	return false
}
