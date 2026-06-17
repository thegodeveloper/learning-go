package concurrency

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewSimpleModule("concurrency", Run))

	registry.Register(registry.NewModule("concurrency", Run, map[string]func(bool){
		"orders": Orders,
	}))
}
