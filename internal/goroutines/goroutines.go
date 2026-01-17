package goroutines

import (
	"fmt"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Goroutines")

		implementCountTo()
		implementCancelFunctionTerminateGoroutine()
		implementBackPressure()
		passingCopyGoroutine()
	}
}

// AllExamples runs all goroutine examples
func AllExamples(show bool) {
	if !show {
		return
	}
	Run(true)
}
