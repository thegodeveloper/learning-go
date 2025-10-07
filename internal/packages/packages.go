package packages

import (
	"fmt"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Packages in Go")
		number1 := 7
		number2 := 17

		// use the add function of calculator package
		fmt.Println(Add(number1, number2))

		// use subtract function of calculator package
		fmt.Println(Subtract(number1, number2))

	}
}
