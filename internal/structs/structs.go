package structs

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("--- Structs")

		definition(false)

		demo(true)

		activity(false)
	}
}
