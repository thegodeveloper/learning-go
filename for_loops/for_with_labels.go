package main

import "fmt"

func forWithLabels() {
	fmt.Println("--- for with labels ---")
	samples := []string{"hello", "world!"}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
	}
}
