package main

import (
	"fmt"
	"strings"
)

// https://www.acmicpc.net/problem/2439
func main() {
	solution1()
}

func solution1() {
	var n int
	_, _ = fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		space := strings.Repeat(" ", n-i)
		star := strings.Repeat("*", i)
		fmt.Println(space + star)
	}
}
