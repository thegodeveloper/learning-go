package wTest

import (
	"fmt"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Tests")
	}
}

func addNumbers(x, y int) int {
	return x + y
}
