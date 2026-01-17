package forloops

import "fmt"

// ForBreakFromSwitch demonstrates breaking out of a for loop from within a switch
func ForBreakFromSwitch(show bool) {
	if show {
		fmt.Println("--- Breaking a for from a switch")

	loop:
		for i := 0; i < 10; i++ {
			switch {
			case i%2 == 0:
				fmt.Println(i, "is even")
			case i%3 == 0:
				fmt.Println(i, "is divisible by 3 but not 2")
			case i%7 == 0:
				fmt.Println("exit the loop!")
				break loop
			default:
				fmt.Println(i, "is boring")
			}
		}
	}
}
