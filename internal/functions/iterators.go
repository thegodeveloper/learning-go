package functions

import "fmt"

func Iterators(show bool) {
	if show {
		fmt.Println("--- Function Iterators ---")

		fmt.Println("First 10 Fibonacci numbers")

		for num := range fibonacci(10) {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}

func fibonacci(limit int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < limit; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}
