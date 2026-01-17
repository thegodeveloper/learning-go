package goroutines

import "fmt"

// ImplementCancelFunctionTerminateGoroutine demonstrates canceling a goroutine
func ImplementCancelFunctionTerminateGoroutine(show bool) {
	if !show {
		return
	}
	implementCancelFunctionTerminateGoroutine()
}

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
