package basics

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Basics in Go")

		net(false)
	}
}
