package concurrency

import "fmt"

func Pipelines(show bool) {
	if show {
		fmt.Println("-- Concurrency Pipelines")

		// Set up the pipeline: gen -> sq -> sq
		for n := range sq(sq(gen(2, 3))) {
			fmt.Println(n)
		}
	}
}

// gen emits the list of ints onto a channel and then closes it
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// sq reads ints from 'in', squares them, and sends them on its returned channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
