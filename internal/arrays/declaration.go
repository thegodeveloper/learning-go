package arrays

import "fmt"

func arrayDeclaration(show bool) {
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

		p := [...]int{1, 2, 3}
		fmt.Printf("%T\n", p)
	}
}
