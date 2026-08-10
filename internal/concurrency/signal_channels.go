package concurrency

import (
	"fmt"
	"time"
)

func signalWorker(done chan struct{}) {
	fmt.Println("Worker: starting work")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker: finished work")
	// Signal completion by sending am empty struct
	done <- struct{}{}
}

func SignalChannels(show bool) {
	if show {
		fmt.Println("--- Signal Channels ---")

		// Channel used only for signaling
		done := make(chan struct{})
		go signalWorker(done)
		// Wait for the signal before exiting
		<-done
		fmt.Println("Main: received signal, exiting")
	}
}
