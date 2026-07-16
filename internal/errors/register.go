package errors

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("errors", Run, map[string]func(bool){
		"basics":          Basics,
		"simpleErrors":    SimpleErrors,
		"sentinelErrors":  SentinelErrors,
		"errorsAreValues": ErrorsAreValues,
		"wrappingErrors":  WrappingErrors,
		"errorsIs":        ErrorsIs,
		"namedErrors":     NamedErrors,
		"allExamples":     AllExamples,
		"customErrorType": CustomErrorType,
		"panicRecover":    PanicRecover,
	}))
}
