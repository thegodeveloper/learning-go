package switches

import "fmt"

func missingLabel(show bool) {
	if show {
		fmt.Println("--- Missing Labels")

		for i := 0; i < 10; i++ {
			switch i {
			case 0, 2, 4, 6:
				fmt.Println(i, "is even")
			case 3:
				fmt.Println(i, "is divisible by 3 but not 2")
			case 7:
				fmt.Println(i, "is the most beautiful number")
			case 8:
				fmt.Println(i, "is so boring, exit the loop")
				break
			default:
				fmt.Println(i, "is either boring")
			}
		}
	}
}
