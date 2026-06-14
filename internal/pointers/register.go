package pointers

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("pointers", Run, map[string]func(bool){
		"declaration": Declaration,
	}))
}
