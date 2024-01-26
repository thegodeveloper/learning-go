package typediffarrayslicesitem

import (
	"fmt"
	"reflect"
)

func Master(show bool) {
	if show {
		a := [5]string{"a", "b", "c", "d", "e"}

		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a[0:3]))
		fmt.Println(reflect.TypeOf(a[0]))
	}
}
