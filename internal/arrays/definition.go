package arrays

import "fmt"

// ArrayDefinition demonstrates array declaration and initialization
func ArrayDefinition(show bool) {
	if show {
		fmt.Println("\n-- Array Declaration")
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

		// you can replace the number that specifies the number of elements in the array with ...
		p := [...]int{1, 2, 3}
		fmt.Printf("%T\n", p)

		// this creates an array of 12 `ints` with the following values:
		// [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
		var x = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)

		// Go has only one-dimensional arrays, but you can simulate multidimensional arrays
		// this declares d to be an array of length 2 whose type is an array of `ints` of length 3
		var d [2][3]int
		d[0] = [3]int{1, 2, 3}
		d[1] = [3]int{4, 5, 6}
		fmt.Println(d)
	}
}
