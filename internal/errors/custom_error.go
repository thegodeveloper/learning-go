package errors

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("values are zero")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func CustomError(show bool) {
	if show {
		fmt.Println("--- Custom Error ---")

		err := check(0, 10)
		if err == nil {
			fmt.Println("check() executed normally!")
		} else {
			fmt.Println(err)
		}

		err = check(0, 0)
		if err.Error() == "values are zero" {
			fmt.Println("Custom error detected!")
		}

		err = formattedError(0, 0)
		if err != nil {
			fmt.Println(err)
		}

		i, err := strconv.Atoi("-123")
		if err == nil {
			fmt.Println("Inv value is:", i)
		}

		i, err = strconv.Atoi("Y123")
		if err != nil {
			fmt.Println(err)
		}
	}
}
