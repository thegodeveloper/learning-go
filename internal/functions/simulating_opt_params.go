package functions

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func simulatingNamedOptionalParams(show bool, opts MyFuncOpts) {
	if show {
		fmt.Println(opts)
	}
}
