package main

import "fmt"

func main() {
	fmt.Println("result : ", solution(70))
}

func solution(angle int) int {
	switch {
	case angle > 0 && angle < 90:
		return 1
	case angle == 90:
		return 2
	case angle > 90 && angle < 180:
		return 3
	case angle == 180:
		return 4
	default:
		return -1
	}
}
