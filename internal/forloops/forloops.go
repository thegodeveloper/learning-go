package forloops

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("---For Loops---")

		Definition(false)
		ForContinue(false)
		ForRange(false)
		ForRangeMaps(false)
		ForString(false)
		ForWithLabels(false)
		ForBreakFromSwitch(false)
		RangeIsACopy(false)
	}
}
