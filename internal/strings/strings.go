package strings

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Strings in Go")

		extractingSingleValue(false)
		toBeAwareInStrings(false)
	}
}
