package regex

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("regex", Run, map[string]func(bool){
		"Demo":               Demo,
		"RegularExpressions": RegularExpressions,
	}))
}
