package chapter6

import "fmt"

func UpdateSlice(s []string, val string) {
	s[len(s)-1] = val
	fmt.Println("in UpdateSlice:", s)
}

func GrowSlice(s []string, val string) {
	s = append(s, val)
	fmt.Println("in GrowSlice:", s)
}

func ex2(show bool) {
	if show {
		s := []string{"a", "b", "c"}
		UpdateSlice(s, "d")
		fmt.Println("in main after UpdateSlice:", s)

		GrowSlice(s, "e")
		fmt.Println("in main aftwer GrowSlice:", s)
	}
}
