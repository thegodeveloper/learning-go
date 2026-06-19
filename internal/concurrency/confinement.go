package concurrency

import "fmt"

func Confinement(show bool) {
	if show {
		fmt.Println("--- Concurrency Confinement")

		// the main goroutine is the consumer, the producer confines the writes
		for v := range producer() {
			fmt.Println(v)
		}
	}
}

func producer() <-chan int {
	out := make(chan int, 7)

	go func() {
		defer close(out)
		for i := 0; i < 7; i++ {
			out <- i
		}
	}()

	return out
}
