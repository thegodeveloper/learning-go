package closures

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("closures", Run, map[string]func(bool){
		"declaration":    Declaration,
		"returnFunction": ReturnFunction,
	}))
}
