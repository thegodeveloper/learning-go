package wArrays

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func Master(show bool) {
	if show {
		fmt.Println("-- Arrays")
		// this creates an array of 5 items, and each item is initialized to 0
		var a [5]int
		fmt.Println(a)
		fmt.Printf("a length: %d\n", len(a))

		// this creates an array of 5 items, but each item is initialized to a specific value
		b := [5]int{0, 1, 2, 3, 4}
		fmt.Println(b)
		fmt.Printf("b length: %d\n", len(b))

		// set the values for the 3 first items
		c := [5]int{0, 1, 2}
		fmt.Println(c)
		fmt.Printf("c length: %d\n", len(c))

		// sparse array: an array where most elements are set to their zero value
		x := [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("x value: ", x)

		arrayDeclaration()

		symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
		fmt.Println(RMB, symbol[RMB])

		// print indexes and elements
		printIndexesAndElements()
	}
}

func arrayDeclaration() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// the elements of a new array variable are initially set to zero value for the element type
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[0])
	fmt.Println(r[2])

	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p)
}

func printIndexesAndElements() {
	fmt.Println("-- Print Indexes and Elements")
	var a [7]int

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}
