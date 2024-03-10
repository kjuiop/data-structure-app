package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var num1, num2 float64

	// 사용자로부터 두 수 입력 받기
	fmt.Print("첫 번째 수를 입력하세요: ")
	if _, err := fmt.Scan(&num1); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	fmt.Print("두 번째 수를 입력하세요: ")
	if _, err := fmt.Scan(&num2); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	result := math.Pow(num1, num2)
	fmt.Println(result)
}
