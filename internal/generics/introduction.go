package generics

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) printAll() {
	if len(s.vals) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Println("Stack elements:")
	for _, element := range s.vals {
		fmt.Print(element)
	}
	fmt.Println("")
}

// Introduction demonstrates basic generics with a Stack example
func Introduction(show bool) {
	if show {
		fmt.Println("--- Introduction to Generics ---")

		var intStack Stack[int]
		intStack.Push(10)
		intStack.Push(20)
		intStack.Push(30)
		intStack.printAll()
		v, ok := intStack.Pop()
		fmt.Println(v, ok)
	}
}
