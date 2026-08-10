package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work.
type Job struct {
	ID int
}

// JResult represents the outcome of processing a job.
type JResult struct {
	JobID  int
	Output string
}

func workerpool(id int, jobs <-chan Job, results chan<- JResult) {
	for job := range jobs {
		// Simulate some work.
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(500 * time.Millisecond)

		// Send the result back.
		results <- JResult{
			JobID:  job.ID,
			Output: fmt.Sprintf("Job %d - worker %d\n", job.ID, id),
		}
	}

	fmt.Printf("Worker %d exiting\n", id)
}

func WorkerPools(show bool) {
	if show {
		fmt.Println("--- Worker Pools ---")

		numWorkers := 3
		numJobs := 20
		jobs := make(chan Job, numJobs)
		results := make(chan JResult, numJobs)

		var wg sync.WaitGroup

		fmt.Printf("%d workers for %d jobs...\n", numWorkers, numJobs)
		for w := 1; w <= numWorkers; w++ {
			workerID := w
			wg.Go(func() {
				workerpool(workerID, jobs, results)
			})
		}

		// Sends jobs to the workers in a separate goroutine.
		wg.Go(func() {
			for j := 1; j <= numJobs; j++ {
				jobs <- Job{ID: j}
			}
			close(jobs)
		})

		// Wait for all workers and close results channel when done.
		go func() {
			wg.Wait()
			close(results)
		}()

		for result := range results {
			fmt.Println("Result: ", result.Output)
		}
		fmt.Println("--- All Jobs Processed ---")
	}
}
