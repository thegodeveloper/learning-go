package functions

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Functions Go")

		SimulatingOptionalParams(false)
		VariadicParameters(false)
		MultipleReturnValues(false)
		NamedReturnValues(false)
		SortSlice(false)
		ReturnFunction(false)
		PrintFileContent(false)
		FunctionsAreValues(false)
		AnonymousFunctions(false)
		FunctionsAsParameters(false)
		Prefixer(false)
	}
}
