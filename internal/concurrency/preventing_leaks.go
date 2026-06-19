package concurrency

import (
	"fmt"
	"time"
)

func PreventingLeaks(show bool) {
	if show {
		fmt.Println("-- Concurrency Preventing Leaks")

		done := make(chan struct{})
		go doWork(done)

		// Cancel after 1 second to avoid leak
		time.Sleep(1 * time.Second)

		close(done)
		// Wait a bit for doWork to print cancellation message
		time.Sleep(100 * time.Millisecond)
	}
}

func doWork(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("doWork canceled")
			return
		default:
		}

		// Simulate work
		fmt.Println("working...")
		time.Sleep(200 * time.Millisecond)
	}
}
