package basics

import "fmt"

func Arithmetic(show bool) {
	if show {
		var a, b = 10, 3
		var result int

		result = a + b
		fmt.Println("Addition:", result)

		result = a - b
		fmt.Println("Subtraction:", result)

		result = a * b
		fmt.Println("Multiplication:", result)

		result = a / b
		fmt.Println("Division:", result)

		result = a % b
		fmt.Println("Modulo:", result)

		const p = 22 / 7.0
		fmt.Println("Decimal value:", p)
	}
}
