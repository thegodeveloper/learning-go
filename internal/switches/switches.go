package switches

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Switches in Go")

		declaration(false)
		switchWithConditions(false)
		switchStatement(false)
		ifElseWithSwitch(false)
		missingLabel(false)
		basicBlankSwitch(false)
	}
}
