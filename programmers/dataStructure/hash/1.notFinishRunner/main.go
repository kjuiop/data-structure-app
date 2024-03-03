package main

import "fmt"

func main() {
	fmt.Println("expected : leo, actual : ", solution([]string{"leo", "kiki", "eden"}, []string{"eden", "kiki"}))
	fmt.Println("expected : vinko, actual : ", solution([]string{"marina", "josipa", "nikola", "vinko", "filipa"}, []string{"josipa", "filipa", "marina", "nikola"}))
	fmt.Println("expected : mislav, actual : ", solution([]string{"mislav", "stanko", "mislav", "ana"}, []string{"stanko", "ana", "mislav"}))
}

func solution(participant []string, completion []string) string {
	return ""
}
