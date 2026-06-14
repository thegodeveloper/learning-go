package channels

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("channels", Run, map[string]func(bool){
		"bufferedChannel":              BufferedChannel,
		"unbufferedChannel":            UnbufferedChannel,
		"readAndWriteToFromChannel":    ReadAndWriteToFromChannel,
		"closedChannel":                ClosedChannel,
		"channelsAsFunctionParameters": ChannelsAsFunctionParameters,
		"definition":                   Definition,
	}))
}
