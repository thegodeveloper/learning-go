package typediffarrayslicesitem

import (
	"fmt"
	"reflect"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Type diff array slices item in Go")

		a := [5]string{"a", "b", "c", "d", "e"}

		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a[0:3]))
		fmt.Println(reflect.TypeOf(a[0]))
	}
}
