package adapter

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("adapter", Run, map[string]func(bool){
		"macExample":     MacExample,
		"windowsExample": WindowsExample,
	}))
}
