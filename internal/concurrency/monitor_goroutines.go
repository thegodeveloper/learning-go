package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type MRequest struct {
	Increment int
	Reply     chan int
}

func counterMonitor(requests <-chan MRequest) {
	count := 0
	for req := range requests {
		count += req.Increment
		req.Reply <- count
	}
	fmt.Println("Monitor exiting")
}

func MonitorGoroutines(show bool) {
	if show {
		fmt.Println("--- Monitor Goroutines ---")

		requests := make(chan MRequest)
		var wg sync.WaitGroup
		numWorkers := 3
		numIncrements := 2

		// Start worker goroutines
		for i := 0; i < numWorkers; i++ {
			wg.Go(func() {
				reply := make(chan int)
				for j := 0; j < numIncrements; j++ {
					requests <- MRequest{1, reply}
					fmt.Printf("Worker %d sees counter = %d\n", i, <-reply)
					time.Sleep(time.Millisecond * 200)
				}
			})
		}

		// Start goroutine separately
		monitorDone := make(chan struct{})
		go func() {
			counterMonitor(requests)
			close(monitorDone) // signal monitor has exited
		}()

		wg.Wait()       // Wait for all workers to finish
		close(requests) // Now it is safe to close requests
		<-monitorDone   // Wait for monitor to exit
		fmt.Println("All workers and monitor finished.")
	}
}
