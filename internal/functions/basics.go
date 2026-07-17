package functions

import (
	"fmt"
	"strings"
)

func Basics(show bool) {
	if show {
		fmt.Println("--- Functions Basics ---")

		// 1. Basic function call
		fmt.Println(greet("Mihalis"))

		// 2. Function returning multiple values
		result, err := divide(10, 2)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Division result:", result)
		}

		// 3. Named return values
		area, perimeter := rectangleStats(5, 3)
		fmt.Println("Rectangle → A: %.2f, P: %.2f\n", area, perimeter)

		// 4. Variadic function
		joined := joinStrings("-", "Go", "is", "fast", "and", "fun")
		fmt.Println("Joined string:", joined)

		// 5.Higher-order function
		double := func(x int) int { return x * 2 }
		fmt.Println("Applying double twice to 5:", applyTwice(double, 5))
	}
}

func greet(name string) string {
	return "Hello " + name + "!"
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func rectangleStats(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // uses named returns
}

func joinStrings(separator string, words ...string) string {
	return strings.Join(words, separator)
}

func applyTwice(fn func(int) int, value int) int {
	return fn(fn(value))
}
