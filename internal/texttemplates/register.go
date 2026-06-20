package texttemplates

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("texttemplates", Run, map[string]func(bool){
		"demo": Demo,
	}))
}
