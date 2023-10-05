package wErrorOld

import (
	"errors"
	"fmt"
	"os"
)

func errorFromFunction(show bool) {
	if show {
		rem, mod, err := calcRemainderAdnMod(7, 7)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(rem, mod)
	}
}

func calcRemainderAdnMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
