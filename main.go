package main

import (
	"github.com/williammunozr/learning-go/channels"
	"github.com/williammunozr/learning-go/goroutines"
	"github.com/williammunozr/learning-go/wAdapter"
	"github.com/williammunozr/learning-go/wArrays"
	"github.com/williammunozr/learning-go/wErrors"
	"github.com/williammunozr/learning-go/wExternalPackages"
	"github.com/williammunozr/learning-go/wInterfaces"
	"github.com/williammunozr/learning-go/wPackages"
	"github.com/williammunozr/learning-go/wRand"
)

func main() {
	// Errors
	wErrors.Master(false)

	// Arrays
	wArrays.Master(false)

	// Interfaces
	wIntefaces.Master(false)

	// Random Values
	wRand.Master(false)

	// Packages
	wPackages.Master(false)

	// External Packages
	wExternalPackages.Master(false)

	// Adapter Pattern
	wAdapter.Master(false)

	// Channels
	channels.Master(false)

	// Goroutines
	goroutines.Master(false)
}
