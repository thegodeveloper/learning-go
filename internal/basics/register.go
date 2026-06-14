package basics

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("basics", Run, map[string]func(bool){
		"net":                Net,
		"arithmetic":         Arithmetic,
		"forLoops":           ForLoops,
		"breakContinueLoops": BreakContinueLoops,
		"outerLoops":         OuterLoops,
		"forRange":           ForRange,
		"forASWhile":         ForASWhile,
	}))
}
