package main

import (
	"github.com/williammunozr/learning-go/wArrays"
	"github.com/williammunozr/learning-go/wErrors"
	"github.com/williammunozr/learning-go/wInterfaces"
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
	wRand.Master(true)
}
