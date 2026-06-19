package concurrency

import "fmt"

func HandyGeneratorsPipelines(show bool) {
	if show {
		fmt.Println("-- Concurrency Handy Generators for Pipelines")

		done := make(chan struct{})
		defer close(done)

		// Repeat 1, 2, 3 until taken 5 values
		for v := range take(done, repeat(done, 1, 2, 3), 5) {
			fmt.Print(v, " ")
		}
	}
}

// repeat sends the given values endlessly on a channel until done is closed
func repeat(done <-chan struct{}, values ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}
	}()

	return out
}

// take reads up to 'n' values from 'in' and then closes
func take(done <-chan struct{}, in <-chan int, n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case out <- <-in:
			}
		}
	}()

	return out
}
