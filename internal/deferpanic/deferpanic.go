package deferpanic

import "fmt"

func closeMsg() {
	fmt.Println("Closed!!!")
}

func Run(show bool) {
	if show {
		fmt.Println("-- Defer Panic in Go")

		defer closeMsg()

		fmt.Println("Doing something...")
		defer fmt.Println("Certainly closed!!!")
		fmt.Println("Doing something else")
	}
}
