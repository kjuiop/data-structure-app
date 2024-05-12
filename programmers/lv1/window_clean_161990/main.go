package main

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/172928
func main() {
	fmt.Println(solution([]string{".#...", "..#..", "...#."}))
	fmt.Println(solution([]string{"..........", ".....#....", "......##..", "...##.....", "....#....."}))
	fmt.Println(solution([]string{".##...##.", "#..#.#..#", "#...#...#", ".#.....#.", "..#...#..", "...#.#...", "....#...."}))
	fmt.Println(solution([]string{"..", "#."}))
	fmt.Println(solution([]string{"#.", ".."}))
}

func solution(wallpaper []string) []int {

	// 빈칸이 . ,  파일이 있으면 #
	// 최소한의 이동거리를 갖는 한 번의 드래그로 모든 파일을 선택해서 한 번에 지우려고 함
	// 반환 값은 드래그의 시작점과 끝점

	// 맨 왼쪽 좌표와 맨 오른쪽 좌표를 구하라
	a, b, c, d := 51, 51, 0, 0
	for idx, val := range wallpaper {

		start := strings.Index(val, "#")
		end := strings.LastIndex(val, "#")

		if start < 0 {
			continue
		}

		if start < b {
			if idx < a {
				a = idx
			}
			b = start
		}

		if end >= d && end >= b {
			d = end + 1
		}

		if c <= idx {
			c = idx + 1
		}
	}

	return []int{a, b, c, d}
}

// 16, 23, 25, 26, 27
