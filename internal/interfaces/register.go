package interfaces

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("interfaces", Run, map[string]func(bool){
		"definition":         Definition,
		"zeroValue":          ZeroValue,
		"emptyInterface":     EmptyInterface,
		"implementInterface": ImplementInterface,
		"implicit":           Implicit,
		"shapes":             Shapes,
	}))
}
