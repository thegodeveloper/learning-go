package recursion

import "fmt"

func Definition(show bool) {
	if show {
		fmt.Println("--- Recursion Definition")

		fmt.Println("factorial(5):", factorial(5))
		fmt.Println("factorial(10):", factorial(10))

		fmt.Println("sumOfDigis(7):", sumOfDigits(7))
		fmt.Println("sumOfDigis(12):", sumOfDigits(12))
		fmt.Println("sumOfDigis(12345):", sumOfDigits(12345))
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

func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}
