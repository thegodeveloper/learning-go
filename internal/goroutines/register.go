package goroutines

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("goroutines", Run, map[string]func(bool){
		"implementCountTo":                          ImplementCountTo,
		"implementCancelFunctionTerminateGoroutine": ImplementCancelFunctionTerminateGoroutine,
		"implementBackPressure":                     ImplementBackPressure,
		"passingCopyGoroutine":                      PassingCopyGoroutine,
		"allExamples":                               AllExamples,
	}))
}
