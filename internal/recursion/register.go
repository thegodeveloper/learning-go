package recursion

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("recursion", Run, map[string]func(bool){
		"definition": Definition,
	}))
}
