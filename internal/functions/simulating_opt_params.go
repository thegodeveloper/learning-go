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

func mainSimulatingOptionalParams(show bool) {
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
