package goroutines

import "fmt"

func implementCountTo() {
	for i := range countTo(7) {
		fmt.Println(i)
	}
}
