package wIntefaces

import "fmt"

func emptyInterface() {
	// empty interface
	// a way to say that a variable could store a value of any type
	var i interface{}
	i = 20
	fmt.Printf("i value: %v, i type: %T\n", i, i)

	i = "hallo"
	fmt.Printf("i value: %v, i type: %T\n", i, i)

	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}
	fmt.Printf("i value: %v, i type: %T\n", i, i)
}
