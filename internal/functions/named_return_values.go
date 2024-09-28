package functions

import (
	"errors"
	"fmt"
)

func divAndRemainderNamed(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, nil
}

func mainNamedReturnValues(show bool) {
	if show {
		fmt.Println("--- Named Return Values ---")

		result, remainder, err := divAndRemainderNamed(7, 7)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("result: %d, remainder: %d\n", result, remainder)
	}
}
