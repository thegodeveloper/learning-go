package adapter

import (
	"fmt"
)

func Run(show bool) {
	if show {
		fmt.Println("-- The Adapter Pattern in Go")
		MacExample(true)
		WindowsExample(true)
	}
}

// MacExample demonstrates using Mac with Lightning connector directly
func MacExample(show bool) {
	if show {
		fmt.Println("\n-- Mac Example (Direct Connection)")
		client := &Client{}
		mac := &Mac{}
		client.InsertLightningConnectorIntoComputer(mac)
	}
}

// WindowsExample demonstrates using Windows with adapter pattern
func WindowsExample(show bool) {
	if show {
		fmt.Println("\n-- Windows Example (Using Adapter)")
		client := &Client{}
		windowsMachine := &Windows{}
		windowsMachineAdapter := &WindowsAdapter{
			windowMachine: windowsMachine,
		}
		client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
	}
}
