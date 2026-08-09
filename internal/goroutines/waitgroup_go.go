package goroutines

import (
	"fmt"
	"sync"
)

func processTask(task string) func() {
	return func() {
		fmt.Println("Processing:", task)
	}
}

func WaitGroupGo(show bool) {
	if show {
		var wg sync.WaitGroup
		fmt.Println("--- WaitGroup.Go Init ---")
		tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7"}

		for _, task := range tasks {
			// wg.Go handles wg.Add(1) and defer wg.Done() internally!
			wg.Go(processTask(task))
		}

		wg.Wait()
		fmt.Println("--- WaitGroup.Go End ---")
	}
}
