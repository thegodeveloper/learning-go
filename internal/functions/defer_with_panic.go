package functions

import "fmt"

func DeferWithPanic(show bool) {
	if show {
		fmt.Println("Defer With Panic")

		processPanic(7)
		processPanic(-7)
	}
}

func processPanic(i int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	if i < 0 {
		fmt.Println("Statement Before Panic")
		panic("input must not be less than zero")
	}

	fmt.Println("Processing input:", i)
}
