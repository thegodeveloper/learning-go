package recursion

import "fmt"

func Definition(show bool) {
	if show {
		fmt.Println("--- Recursion Definition")

		fmt.Println("factorial(5):", factorial(5))
		fmt.Println("factorial(10):", factorial(10))
	}
}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// n * (n - 1) * (n - 2) * factorial (n - 3)... until factorial(0) = 1
	return n * factorial(n-1)
}
