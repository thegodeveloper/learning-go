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
