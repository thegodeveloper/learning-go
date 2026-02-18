package generics

import (
	"errors"
	"fmt"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func GenericsDemo(show bool) {
	if show {
		fmt.Println("Generics Demo")

		var a uint = 18_446_744_073_709_551_615
		var b uint = 9_223_372_036_854_775_808

		fmt.Println(DivRemainder(a, b))
	}
}

func DivRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}
