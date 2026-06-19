package concurrency

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("concurrency", Run, map[string]func(bool){
		"orders":           Orders,
		"confinement":      Confinement,
		"forselectloop":    ForSelectLoop,
		"preventingleaks":  PreventingLeaks,
		"orchannelpattern": OrChannelPattern,
		"errorhandling":    ErrorHandling,
		"pipelines":        Pipelines,
	}))
}
