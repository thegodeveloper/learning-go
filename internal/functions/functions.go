package functions

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Functions Go")

		mainSimulatingOptionalParams(false)
		mainVariadicParameters(false)
		mainMultipleReturnValues(false)
		mainNamedReturnValues(false)
		mainSortSlice(false)
		mainReturnFunction(false)
		mainPrintFileContent(false)
		mainFunctionsAreValues(false)
		mainAnonymousFunctions(false)
		mainFunctionsAsParameters(false)
		mainPrefixer(false)
	}
}
