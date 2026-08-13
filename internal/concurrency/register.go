package concurrency

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("concurrency", Run, map[string]func(bool){
		"orders":                   Orders,
		"confinement":              Confinement,
		"forselectloop":            ForSelectLoop,
		"preventingleaks":          PreventingLeaks,
		"orchannelpattern":         OrChannelPattern,
		"errorhandling":            ErrorHandling,
		"pipelines":                Pipelines,
		"selectcon":                SelectCon,
		"sendreceive":              SendReceive,
		"BufferedUnbuffered":       BufferedUnbuffered,
		"SelectTimeOut":            SelectTimeOut,
		"WorkerPools":              WorkerPools,
		"SignalChannels":           SignalChannels,
		"MonitorGoroutines":        MonitorGoroutines,
		"Matrix":                   Matrix,
		"handygeneratorspipelines": HandyGeneratorsPipelines,
	}))
}
