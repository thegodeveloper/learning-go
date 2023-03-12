package main

import (
	"Packages/calculator"
	"fmt"
)

func main() {
	number1 := 7
	number2 := 17

	// use the add function of calculator package
	fmt.Println(calculator.Add(number1, number2))

	// use subtract function of calculator package
	fmt.Println(calculator.Subtract(number1, number2))

}
