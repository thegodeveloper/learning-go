package rand

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("rand", Run, map[string]func(bool){
		"Demo":         Run,
		"RandomNumber": RandomNumber,
	}))
}
