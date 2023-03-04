package main

import "fmt"

// pointers in Go are types that store the address of the value, not the value.
// A pointer is simply a variable whose contents are the address where another variable is stored.

// you cannot store int in intPtr, you can only store the address of int
var intPtr *int

func a(i int) {
	i = 0
}

func b(i *int) {
	*i = 101
}

func main() {
	x := 100

	a(x)
	fmt.Println(x)

	b(&x)
	fmt.Println(x)

	fmt.Println(&x)

	// assign to intPtr the address of our x variable
	intPtr = &x
	fmt.Printf("x address: %v\n", &x)
	fmt.Printf("intPtr address: %v\n", &intPtr)
	fmt.Printf("intPtr value: %v\n", *intPtr)
	fmt.Printf("x value: %v\n", x)
	fmt.Printf("intPtr value: %v\n", *intPtr)
	*intPtr = 80
	fmt.Printf("x value: %v\n", x)

	// pointers in functions
	say := "hello"
	changeValue(&say) // pass a pointer
	fmt.Printf("now say variable has the value: %v\n", say)
}

func changeValue(word *string) {
	// add "world" to the string pointed to by 'word'
	*word += "world"
}
