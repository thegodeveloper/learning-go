package packages

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewSimpleModule("packages", Run))
}

func init() {
	registry.Register(registry.NewModule("packages", Run, map[string]func(bool){
		"ConsumePackage": ConsumePackage,
	}))
}
