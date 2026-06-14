package sumtwonumbers

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewSimpleModule("sumtwonumbers", Run))
}
