package functions

import "fmt"

func RecoverFunction(show bool) {
	if show {
		fmt.Println("Recover Function")

		processRecover()
		fmt.Println("Returned from processRecover function")
	}
}

func processRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process...")
	panic("Something went wrong")
}
