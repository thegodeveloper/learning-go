package concurrency

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Concurrency")

		implementation()
	}
}

func implementation() {
	result, err := timeLimit()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {
	time.Sleep(3 * time.Second)
	return 7, nil
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	fmt.Println("-- Getting into SlowComplicatedParser")
	return parser
}
