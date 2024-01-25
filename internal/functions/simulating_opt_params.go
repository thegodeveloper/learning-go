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
