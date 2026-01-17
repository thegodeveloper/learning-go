package switches

import "fmt"

func Declaration(show bool) {
	if show {
		var finger = 1

		switch finger {
		case 0:
			fmt.Println("Thumb")
		case 1:
			fmt.Println("Index")
		case 2:
			fmt.Println("Middle")
		case 3:
			fmt.Println("Ring")
		case 4:
			fmt.Println("Pinkie")
		default:
			fmt.Println("Humans usually have no more than five fingers")
		}
	}
}
