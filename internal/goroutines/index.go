package goroutines

import (
	"fmt"
)

func Index(show bool) {
	if show {
		fmt.Println("-- Goroutines")

		implementCountTo()
		implementCancelFunctionTerminateGoroutine()
		implementBackPressure()
		passingCopyGoroutine()
	}
}
