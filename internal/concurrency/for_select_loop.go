package concurrency

import (
	"fmt"
	"time"
)

func ForSelectLoop(show bool) {
	if show {
		fmt.Println("--- Concurrency for-select loop")

		done := make(chan struct{})
		go worker(done)
		time.Sleep(2 * time.Second)

		// signal working to exit
		close(done)
	}
}

func worker(done <-chan struct{}) {
	for {
		select {
		case <-done:
			// Stop the loop and exit
			return
		default:
		}

		// Perform non-blocking work
		fmt.Println("working...")
		time.Sleep(100 * time.Millisecond)
	}
}
