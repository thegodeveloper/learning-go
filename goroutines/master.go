package goroutines

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("-- Goroutines")
		cancelFunctionToTerminateGoroutine()
	}
}

func cancelFunctionToTerminateGoroutine() {

}

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				retunr
			}
		}
	}
}
