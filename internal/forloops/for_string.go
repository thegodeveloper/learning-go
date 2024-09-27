package forloops

import "fmt"

func forString(show bool) {
	if show {
		fmt.Println("--- for in string")

		samples := []string{"hello", "world!"}
		for _, sample := range samples {
			for i, r := range sample {
				fmt.Println(i, r, string(r))
			}
		}
	}
}
