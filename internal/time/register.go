package time

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("time", Run, map[string]func(bool){
		"Demo":          Demo,
		"SpecificTime":  SpecificTime,
		"TransformTime": TransformTime,
		"Formatting":    Formatting,
	}))
}
