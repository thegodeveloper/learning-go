package main

import "fmt"

func main() {
	// declare a variable of type int with a value of 10.
	count := 10

	// display the "value of" and "address of" count.
	fmt.Println("count:\tValue Of[", count, "]\tAddress Of[", &count, "]")

	// pass the "value of" the count.
	increment(count)

	// display the "value of" and "address of" count.
	fmt.Println("count:\tValue Of[", count, "]\tAddress Of[", &count, "]")
}

//go:noinline
func increment(inc int) {
	// increment the "value of" inc.
	inc += 1
	fmt.Println("inc:\t Value Of[", inc, "]\tAddr Of[", &inc, "]")
}
