package wcomplex

import (
	"fmt"
	"math/cmplx"
)

func Master(show bool) {
	if show {
		var a = complex(6, 2) // type can be omitted
		var b complex64 = complex(9, 2)
		var c = complex(7.7, 7.7)
		fmt.Println("complex128 a = ", a)
		fmt.Println("complex64 b = ", b)
		fmt.Println("complex128 c = ", c)

		fmt.Println("a + c = ", a+c)
		fmt.Println("a - c = ", a, c)
		fmt.Println("a * c = ", a, c)
		fmt.Println("a / c = ", a, c)
		fmt.Println("real(a) = ", real(a))
		fmt.Println("imag(a) = ", imag(a))
		fmt.Println("Abs(a) = ", cmplx.Abs(a))

		fmt.Printf("The type of a is %T and the type of b is %T", a, b)
	}
}
