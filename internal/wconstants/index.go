package wconstants

import (
	"fmt"
	"reflect"
)

// untyped constant
const x = 10

// typed constant
const i int = 10

const (
	Pi               = 3.14
	Avogadro float32 = 6.022e23
)

func Index(show bool) {
	if show {
		fmt.Println("What is the value of Pi? Pi is", Pi)
		fmt.Println(reflect.TypeOf(Pi))
		fmt.Println("Avogadro's Number value is", Avogadro)
		fmt.Println(reflect.TypeOf(Avogadro))

		// valid declarations
		var y = x
		var z float64 = x
		var d byte = x

		i := 20
		f := float64(i)
		fmt.Println(f)

		fmt.Println("value of y: ", y)
		fmt.Println("value of z: ", z)
		fmt.Println("value of d: ", d)

		fmt.Println("value of i: ", i)
	}
}
