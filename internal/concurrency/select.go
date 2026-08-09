package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func workerSelect(name string, ch chan<- string) {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)
	ch <- fmt.Sprintf("%s finished after %v", name, delay)
}

func SelectCon(show bool) {
	if show {
		fmt.Println("--- Select Concurrency ---")

		ch1 := make(chan string)
		ch2 := make(chan string)
		go workerSelect("Worker 1", ch1)
		go workerSelect("Worker 2", ch2)

		for range 2 {
			select {
			case msg1 := <-ch1:
				fmt.Println("Received:", msg1)
			case msg2 := <-ch2:
				fmt.Println("Received:", msg2)
			}
		}
		fmt.Println("All workers reported back!")
	}
}
