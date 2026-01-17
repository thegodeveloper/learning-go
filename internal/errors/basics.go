package errors

import (
	"errors"
	"fmt"
	"os"
)

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

// Basics demonstrates basic error handling with division
func Basics(show bool) {
	if !show {
		return
	}
	basics()
}

func basics() {
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
