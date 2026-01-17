package slices

import (
	"fmt"
	"slices"
)

func CompareSlices(show bool) {
	if show {
		fmt.Println("--- Compare Slices")

		x := []int{1, 2, 3, 4, 5}
		y := []int{1, 2, 3, 4, 5}
		z := []int{1, 2, 3, 4, 5, 6}
		s := []string{"a", "b", "c"}

		fmt.Println("slices x and y are equal: ", slices.Equal(x, y))
		fmt.Println("slices x and z are equal: ", slices.Equal(x, z))
		fmt.Printf("Slices x and s cannot be compare because x is of type %T and s is of type %T", x, s)
		//fmt.Println("slices x and s are equal: ", slices.Equal(x, s))
	}
}
