package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func Demo(show bool) {
	if show {
		fmt.Println("--- Goroutines Demo ---")

		numGoroutines := 7

		fmt.Printf("Starting %d goroutines...\n", numGoroutines)
		var wg sync.WaitGroup

		// Launch goroutines without any synchronization.
		for i := range numGoroutines {
			wg.Add(1)
			go worker(i, &wg)
		}

		wg.Wait()
	}

	fmt.Println("--- Goroutines Demo function is exiting without waiting! ---")
}
