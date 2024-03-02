package goroutines

import (
	"fmt"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Goroutines")

		implementCountTo()
		implementCancelFunctionTerminateGoroutine()
		implementBackPressure()
		passingCopyGoroutine()
	}
}
