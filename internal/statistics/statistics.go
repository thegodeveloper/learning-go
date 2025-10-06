package statistics

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Run(show bool) {
	if show {
		arguments := os.Args
		if len(arguments) == 1 {
			fmt.Println("Usage: go run ./cmd/cli [options]")
			return
		}

		var minv, maxv float64
		var initialized = 0
		nValues := 0
		var sum float64
		for i := 1; i < len(arguments); i++ {
			n, err := strconv.ParseFloat(arguments[i], 64)
			if err != nil {
				continue
			}
			nValues++
			sum += n
			if initialized == 0 {
				minv = n
				maxv = n
				initialized = 1
				continue
			}
			if n < minv {
				minv = n
			}
			if n > maxv {
				maxv = n
			}
		}
		fmt.Println("Number of values:", nValues)
		fmt.Println("Minimum value:", minv)
		fmt.Println("Maximum value:", maxv)

		if nValues == 0 {
			return
		}
		meanValue := sum / float64(nValues)
		fmt.Printf("Mean value: %.5f\n", meanValue)
		// standard deviation
		var squared float64
		for i := 1; i < len(arguments); i++ {
			n, err := strconv.ParseFloat(arguments[i], 64)
			if err != nil {
				continue
			}
			squared += math.Pow(n-meanValue, 2)
		}
		standardDeviation := math.Sqrt(squared / float64(nValues))
		fmt.Printf("Standard Deviation: %.5f\n", standardDeviation)
	}
}
