// Package calculator does simple calculations
package calculator

import "fmt"

// Add takes two numbers and returns the result of adding
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and returns the result of subtracting b from a
func Subtract(a, b float64) float64 {
	return a - b
}

func Run(show bool) {
	if show {
		fmt.Println("-- Calculator App")

		fmt.Println("7 + 7:", Add(7.0, 7.0))
		fmt.Println("7 - 7:", Subtract(7.0, 7.0))
	}
}
