package main

import "fmt"

func main() {
	fmt.Println("result : ", solution(10))
}

func solution(n int) int {

	if n < 0 || n > 1000 {
		return -1
	}

	result := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result += i
		}
	}

	return result
}
