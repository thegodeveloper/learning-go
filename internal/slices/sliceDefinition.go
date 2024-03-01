package slices

import "fmt"

func sliceDefinition(show bool) {
	if show {
		fmt.Println("--- Slice Definition")

		// this is a string array of 5 elements
		a := [5]string{"a", "b", "c", "d", "e"}

		fmt.Println(a)
		fmt.Println(a[:])
		fmt.Println(a[0])
		fmt.Println(a[0], a[1], a[2], a[3], a[4])
		fmt.Println(a[0:2])
		fmt.Println(a[1:4])
		fmt.Println(a[:2])
		fmt.Println(a[2:])
	}
}
