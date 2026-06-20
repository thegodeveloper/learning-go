package interfaces

import "fmt"

type zeroValue interface{}

// ZeroValue demonstrates zero value of interfaces
func ZeroValue(show bool) {
	if !show {
		return
	}

	fmt.Println("--- Zero Value Interfaces")

	zerovalue()
}

func zerovalue() {
	var a zeroValue

	fmt.Println(a) // <nil>
}
