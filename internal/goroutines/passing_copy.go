package goroutines

import (
	"fmt"
	"sync"
)

// PassingCopyGoroutine demonstrates passing values to goroutines
func PassingCopyGoroutine(show bool) {
	if !show {
		return
	}
	passingCopyGoroutine()
}

func passingCopyGoroutine() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}

	wg.Wait()
}
