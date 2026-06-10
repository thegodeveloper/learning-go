package generics

import "fmt"

type ValidTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Doubler[T ValidTypes](t T) T {
	return t * 2
}

func Definition(show bool) {
	if show {
		fmt.Println("Generics Definition")

		fmt.Println(Doubler(7))
		fmt.Println(Doubler(7.7))
	}
}
