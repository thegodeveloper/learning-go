package main

import "fmt"

func forString() {
	samples := []string{"hello", "world!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
}
