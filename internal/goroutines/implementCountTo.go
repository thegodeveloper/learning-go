package goroutines

import "fmt"

// ImplementCountTo demonstrates using a channel to count
func ImplementCountTo(show bool) {
	if !show {
		return
	}
	implementCountTo()
}

func implementCountTo() {
	for i := range countTo(7) {
		fmt.Println(i)
	}
}
