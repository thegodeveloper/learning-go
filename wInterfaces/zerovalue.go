package wIntefaces

import "fmt"

type zeroValue interface{}

func zerovalue() {
	var a zeroValue

	fmt.Println(a) // <nil>
}
