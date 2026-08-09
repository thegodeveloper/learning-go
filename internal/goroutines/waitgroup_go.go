package goroutines

import (
	"fmt"
	"sync"
)

func WaitGroupGo(show bool) {
	if show {
		var wg sync.WaitGroup
		fmt.Println("--- WaitGroup.Go Init ---")
		tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7"}

		for _, task := range tasks {
			wg.Go(func() {
				fmt.Println("Processing:", task)
			})
		}
		wg.Wait()
		fmt.Println("--- WaitGroup.Go End ---")
	}
}
