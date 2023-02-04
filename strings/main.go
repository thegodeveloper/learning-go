package main

import "fmt"

func main() {
	extractingSingleValue()
}

func extractingSingleValue() {
	fmt.Println("--- extracting a single value from a string ---")
	s := "Hello Golang!"
	var b byte = s[7]
	fmt.Println("s:", s)
	fmt.Println("b:", b)
}
