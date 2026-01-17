package strings

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Strings in Go")

		ExtractingSingleValue(false)
		ToBeAwareInStrings(false)
	}
}
