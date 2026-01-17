package maps

import "fmt"

func MapSet(show bool) {
	if show {
		fmt.Println("--- Map Set")

		intSet := map[int]bool{}
		vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
		for _, v := range vals {
			intSet[v] = true
		}
		fmt.Println(len(vals), len(intSet))
		fmt.Println(intSet[5])
		fmt.Println(intSet[500])
		if intSet[100] {
			fmt.Println("100 is in the set")
		}
	}
}
