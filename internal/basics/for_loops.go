package basics

import (
	"fmt"
)

func ForLoops(show bool) {
	// simple iteration over a range
	for i := 0; i < 10; i++ {
		fmt.Println("For loop:", i)
	}

	// iterate over a collection
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range numbers {
		fmt.Printf("index: %d, Value: %d\n", index, value)
	}
}
