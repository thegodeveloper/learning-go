// Package calculator does simple calculations
package calculator

import "fmt"

// Add takes two numbers and returns the result of adding
func Add(a, b float64) float64 {
	return a + b
}

func Index(show bool) {
	if show {
		fmt.Println("Calculator App")
	}
}
