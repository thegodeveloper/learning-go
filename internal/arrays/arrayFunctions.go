package arrays

import (
	"fmt"
)

// ArrayFunctions demonstrates array functions like len
func ArrayFunctions(show bool) {
	if show {
		fmt.Println("--- Array Functions")

		a := [...]int{0, 1, 2, 3, 4, 5, 6}

		fmt.Println("a len: ", len(a))
	}
}
