package forloops

import "fmt"

func definition(show bool) {
	if show {
		fmt.Println("--- Definition")
		x := 5
		counter := x

		// for with a condition
		for counter > 0 {
			fmt.Println(counter)
			counter--
		}

		// standard for loop
		for i := 0; i < x; i++ {
			fmt.Print(i)
		}
		fmt.Println()

		// for with continue and break, while loop simulation
		for {
			if x%2 != 0 {
				fmt.Printf("%d is odd\n", x)
				x++
				continue
			}
			break
		}

		// infinite for loop with break
		for {
			fmt.Println("Never stop")
			break
		}

		// idiomatic for loop
		i := 0
		for ok := true; ok; ok = i != 10 {
			fmt.Print(i*i, " ")
			i++
		}

		// for with range
		aSlice := []int{-1, 2, 1, -1, 2, -2}
		for i, v := range aSlice {
			fmt.Println("index:", i, "value:", v)
		}
	}
}
