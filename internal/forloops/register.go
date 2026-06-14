package forloops

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("forloops", Run, map[string]func(bool){
		"definition":         Definition,
		"forContinue":        ForContinue,
		"forRange":           ForRange,
		"forRangeMaps":       ForRangeMaps,
		"forString":          ForString,
		"forWithLabels":      ForWithLabels,
		"forBreakFromSwitch": ForBreakFromSwitch,
		"rangeIsACopy":       RangeIsACopy,
	}))
}
