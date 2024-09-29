package functions

import "fmt"

func mainAnonymousFunctions(show bool) {
	if show {
		fmt.Println("--- Anonymous Functions ---")

		f := func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}

		for i := 0; i < 5; i++ {
			f(i)
		}
	}
}
