package forloops

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("---For Loops---")

		definition(false)
		forContinue(false)
		forRange(false)
		forRangeMaps(false)
		forString(false)
		forWithLabels(false)
		forBreakFromSwitch(false)
	}
}
