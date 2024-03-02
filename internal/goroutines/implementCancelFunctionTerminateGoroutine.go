package goroutines

import "fmt"

func implementCancelFunctionTerminateGoroutine() {
	fmt.Println("-- Implement Cancel Function Terminate Goroutine")
	ch, cancel := countTo2(7)
	for i := range ch {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}
