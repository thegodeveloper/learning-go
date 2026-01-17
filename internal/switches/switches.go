package switches

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Switches in Go")

		Declaration(false)
		SwitchWithConditions(false)
		SwitchStatement(false)
		IfElseWithSwitch(false)
		MissingLabel(false)
		BasicBlankSwitch(false)
	}
}
