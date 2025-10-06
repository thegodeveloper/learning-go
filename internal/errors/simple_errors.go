package errors

import (
	"errors"
	"fmt"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func simpleErrors() {
	evenNumber, err := doubleEven(7)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	fmt.Printf("%d is even", evenNumber)
}
