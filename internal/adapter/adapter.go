package adapter

import (
	"fmt"
)

func Run(show bool) {
	if show {
		fmt.Println("-- The Adapter Pattern in Go")
		client := &Client{}
		mac := &Mac{}

		client.InsertLightningConnectorIntoComputer(mac)

		windowsMachine := &Windows{}
		windowsMachineAdapter := &WindowsAdapter{
			windowMachine: windowsMachine,
		}

		client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
	}
}
