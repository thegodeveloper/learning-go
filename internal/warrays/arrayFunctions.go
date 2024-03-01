package warrays

import (
	"fmt"
)

func arrayFunctions(show bool) {
	if show {
		fmt.Println("--- Array Functions")

		a := [...]int{0, 1, 2, 3, 4, 5, 6}

		fmt.Println("a len: ", len(a))
	}
}
