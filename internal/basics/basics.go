package basics

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Basics in Go")

		Net(false)

		Arithmetic(false)

		ForLoops(false)

		BreakContinueLoops(false)

		OuterLoops(false)

		ForRange(false)
	}
}
