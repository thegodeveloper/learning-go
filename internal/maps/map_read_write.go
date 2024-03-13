package maps

import "fmt"

func mapReadWrite(show bool) {
	if show {
		fmt.Println("--- Map Read Write")

		totalWins := map[string]int{}
		totalWins["Orcas"] = 1
		totalWins["Lions"] = 2
		fmt.Println("Orcas:", totalWins["Orcas"])
		fmt.Println("Lions:", totalWins["Lions"])
		fmt.Println("Kittens:", totalWins["Kittens"])
		totalWins["Kittens"]++
		fmt.Println("Kittens:", totalWins["Kittens"])
		totalWins["Lions"] = 3
		fmt.Println("Lions:", totalWins["Lions"])
	}
}
