package functions

import (
	"errors"
	"fmt"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

// MultipleReturnValues demonstrates returning multiple values from functions
func MultipleReturnValues(show bool) {
	if show {
		fmt.Println("--- Multiple return values ---")

		// Returning multiple values
		result, remainder, err := divAndRemainder(7, 7)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("result: %d, remainder: %d", result, remainder)
	}
}
