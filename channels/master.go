package channels

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("-- Channels")
		definition()

		fmt.Println("-- Channels as Function Parameters")
		channelsAsFunctionParameters()
	}
}
