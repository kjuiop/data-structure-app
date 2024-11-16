package main

import (
	"fmt"
	"strings"
)

// https://www.acmicpc.net/problem/2438
func main() {
	solution2()
}

func solution() {
	var n int
	_, _ = fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		var star string
		for j := 1; j <= i; j++ {
			star += "*"
		}
		fmt.Println(star)
	}
}

func solution2() {
	var n int
	_, _ = fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
