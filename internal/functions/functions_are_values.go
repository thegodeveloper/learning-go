package functions

import "fmt"

// FunctionsAreValues demonstrates that functions are first-class values
func FunctionsAreValues(show bool) {
	if show {
		fmt.Println("--- Functions Are Values ---")

		var myFuncVariable func(string) int

		myFuncVariable = f1
		result := myFuncVariable("Hello")
		fmt.Println(result)

		myFuncVariable = f2
		result = myFuncVariable("Hello")
		fmt.Println(result)
	}
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}
