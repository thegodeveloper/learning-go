package functions

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func simulatingNamedOptionalParams(opts MyFuncOpts) {
	fmt.Println(opts)
}

// SimulatingOptionalParams demonstrates simulating optional parameters
func SimulatingOptionalParams(show bool) {
	if show {
		fmt.Println("--- simulating named and optional parameters ---")

		simulatingNamedOptionalParams(MyFuncOpts{
			LastName: "Patel",
			Age:      50,
		})
		simulatingNamedOptionalParams(MyFuncOpts{
			FirstName: "Joe",
			LastName:  "Smith",
		})
	}
}
