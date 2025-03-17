package slices

import "fmt"

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlices(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func passMapSlice(show bool) {
	if show {
		m := map[int]string{
			1: "first",
			2: "second",
		}
		modMap(m)
		fmt.Println(m)

		s := []int{1, 2, 3}
		modSlices(s)
		fmt.Println(s)
	}
}
