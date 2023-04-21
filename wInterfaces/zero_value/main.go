package main

import "fmt"

type zeroValue interface{}

func main() {
	var a zeroValue

	fmt.Println(a) // <nil>
}
