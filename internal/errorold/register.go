package errorold

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("errorold", Run, map[string]func(bool){
		"demo":              Demo,
		"errorsAndValues":   ErrorsAndValues,
		"errorFromFunction": ErrorFromFunction,
		"wrappingErrors":    WrappingErrors,
		"simple":            Simple,
		"sentinelError":     SentinelError,
	}))
}
