package closures

import "fmt"

func ReturnFunction(show bool) {
	if show {
		fmt.Println("--- Return Function ---")

		sequence := adder()

		fmt.Println("the value returned by the closure:", sequence())
	}
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("adding 1 to i:", i)
		return i
	}
}
