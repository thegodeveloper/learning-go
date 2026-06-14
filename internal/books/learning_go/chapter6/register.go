package chapter6

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewSimpleModule("chapter6", Run))
}
