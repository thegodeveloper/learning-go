package channels

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Channels")
		definition()

		fmt.Println("-- Channels as Function Parameters")
		channelsAsFunctionParameters()
	}
}

// Definition demonstrates all basic channel operations
func Definition(show bool) {
	if !show {
		return
	}
	definition()
}
