package forLoops

import "fmt"

func forRange() {
	fmt.Println("--- for range ---")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}
