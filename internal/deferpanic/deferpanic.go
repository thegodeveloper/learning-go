package deferpanic

import "fmt"

func closeMsg() {
	fmt.Println("Closed!!!")
}

func Master(show bool) {
	if show {
		defer closeMsg()

		fmt.Println("Doing something...")
		defer fmt.Println("Certainly closed!!!")
		fmt.Println("Doing something else")
	}
}
