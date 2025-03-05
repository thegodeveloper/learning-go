package functions

import "fmt"

func mainAnonymousFunctions(show bool) {
	if show {
		anonymousFunctionOption1(false)
		anonymousFunctionOption2(false)
	}
}

func anonymousFunctionOption1(show bool) {
	if show {
		fmt.Println("--- Anonymous Functions ---")

		f := func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}

		for i := range 5 {
			f(i)
		}
	}
}

func anonymousFunctionOption2(show bool) {
	if show {
		for i := range 5 {
			func(j int) {
				fmt.Println("printing", j, "from inside of an anonymous function")
			}(i)
		}
	}
}
