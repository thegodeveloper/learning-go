package functions

import "fmt"

func DeferFunctions(show bool) {
	if show {
		fmt.Println("Defer Functions Demonstration")

		process()
	}
}

func process() {
	// the defer statements are executed before the function ends
	// they are executed from lastest defer until first defer
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")

	fmt.Println("Normal statement executed")
}
