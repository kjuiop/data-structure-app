package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := sumPrime(a, b)
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
