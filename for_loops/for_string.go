package main

import "fmt"

func forString() {
	fmt.Println("--- for in string ---")
	samples := []string{"hello", "world!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
}
