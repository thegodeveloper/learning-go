package indirectmemoryaccess

import "fmt"

func Master(show bool) {
	if show {
		// declare variable of type int with a value of 10.
		count := 10

		// display the "value of" and "address of" count.
		fmt.Println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

		// pass the "address of" count.
		increment(&count)

		// display the "value of" and "address of" count after calling increment function.
		fmt.Println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
	}
}

func increment(inc *int) {
	// increment the "value of" count that the "pointer points to". (dereferencing)
	*inc += 1

	// display the "value of" and "address of" inc.
	fmt.Println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
