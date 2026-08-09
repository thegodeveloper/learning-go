package goroutines

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("goroutines", Run, map[string]func(bool){
		"demo":                  Demo,
		"implementCountTo":      ImplementCountTo,
		"implementBackPressure": ImplementBackPressure,
		"passingCopyGoroutine":  PassingCopyGoroutine,
		"allExamples":           AllExamples,
		"waitgroupGo":           WaitGroupGo,
		"implementCancelFunctionTerminateGoroutine": ImplementCancelFunctionTerminateGoroutine,
	}))
}
