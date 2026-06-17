package generics

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("generics", Run, map[string]func(bool){
		"introduction":            Introduction,
		"definition":              Definition,
		"divRemainder":            DemoGenerics,
		"functionsDataStructures": FunctionsWithDataStructures,
	}))
}
