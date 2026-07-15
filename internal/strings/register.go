package strings

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("strings", Run, map[string]func(bool){
		"extractingSingleValue": ExtractingSingleValue,
		"toBeAwareInStrings":    ToBeAwareInStrings,
		"builderString":         BuilderString,
		"stringsRunes":          StringsRunes,
	}))
}
