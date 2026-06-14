package switches

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("switches", Run, map[string]func(bool){
		"declaration":          Declaration,
		"switchWithConditions": SwitchWithConditions,
		"switchStatement":      SwitchStatement,
		"ifElseWithSwitch":     IfElseWithSwitch,
		"missingLabel":         MissingLabel,
		"basicBlankSwitch":     BasicBlankSwitch,
	}))
}
