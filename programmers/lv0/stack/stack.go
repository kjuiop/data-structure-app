package main

import "fmt"

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	stack := make(Stack, 0)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop()) // Should print nil
}
