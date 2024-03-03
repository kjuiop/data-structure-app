package main

import (
	"fmt"
	"log"
	"math"
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

	sum := sumPrime(num1, num2)
	fmt.Println(sum)
}

func sumPrime(a, b int) int {
	sum := 0

	for i := a; i <= b; i++ {
		if i == 1 {
			continue
		}
		if isPrime(i) {
			sum += i
		}
	}

	return sum
}

func isPrime(a int) bool {
	if a <= 1 {
		return false
	}

	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}

func isSuperPrime(a int) bool {
	if a <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}
