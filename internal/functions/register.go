package functions

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("functions", Run, map[string]func(bool){
		"basics":                   Basics,
		"simulatingOptionalParams": SimulatingOptionalParams,
		"variadicParameters":       VariadicParameters,
		"multipleReturnValues":     MultipleReturnValues,
		"namedReturnValues":        NamedReturnValues,
		"sortSlice":                SortSlice,
		"returnFunction":           ReturnFunction,
		"printFileContent":         PrintFileContent,
		"functionsAreValues":       FunctionsAreValues,
		"anonymousFunctions":       AnonymousFunctions,
		"functionsAsParameters":    FunctionsAsParameters,
		"prefixer":                 Prefixer,
		"deferFunctions":           DeferFunctions,
		"deferWithPanic":           DeferWithPanic,
		"recoverFunction":          RecoverFunction,
		"iterators":                Iterators,
	}))
}
