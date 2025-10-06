package interfaces

import "fmt"

func Index(show bool) {
	if show {
		fmt.Println("-- Interface Definition")
		definition()

		fmt.Println("-- Zero Value Interfaces")
		zerovalue()

		fmt.Println("-- Empty Interfaces")
		emptyInterface()

		fmt.Println("-- Implement Interface")
		implementInterface()

		implicit(false)
	}
}
