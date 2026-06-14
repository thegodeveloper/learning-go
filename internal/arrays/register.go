package arrays

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("arrays", Run, map[string]func(bool){
		"definition":              ArrayDefinition,
		"basics":                  Basics,
		"symbols":                 Symbols,
		"printIndexesAndElements": PrintIndexesAndElements,
		"printType":               PrintType,
		"comparingArrays":         ComparingArrays,
		"arrayFunctions":          ArrayFunctions,
		"arrayStrings":            ArrayStrings,
	}))
}
