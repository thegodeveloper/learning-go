package goroutines

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("-- Goroutines")
		cancelFunctionToTerminateGoroutine()
	}
}

func cancelFunctionToTerminateGoroutine() {
	for i := range countTo(7) {
		fmt.Println(i)
	}
}

// Don't do this in a real program
// This violates one of our "when to use concurrency" guidelines
func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
