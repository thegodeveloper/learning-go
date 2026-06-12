package maps

import "fmt"

func MapLiteral(show bool) {
	if show {
		fmt.Println("Map Literal")

		stack := map[string]int{
			"Linux":      1,
			"Kubernetes": 2,
			"Go":         3,
			"AWS":        4,
			"Terraform":  5,
			"Docker":     6,
			"Rust":       7,
		}
		for k, v := range stack {
			fmt.Println(k, "Value:", v)
		}
	}
}
