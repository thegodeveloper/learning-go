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

		// overflow
		var maxInt = 9223372036854775807
		fmt.Println("Max value:", maxInt)
		// it goes to a negative value
		maxInt++
		fmt.Println("Max value:", maxInt)

		// overflow with unsigned integers
		var uMaxInt uint64 = 18446744073709551615
		uMaxInt++
		fmt.Println("Max value uint64:", uMaxInt)
	}
}
