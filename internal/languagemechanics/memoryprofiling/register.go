package memoryprofiling

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewSimpleModule("memoryprofiling", Run))
}
