package arrays

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

// Symbols demonstrates using arrays with custom types
func Symbols(show bool) {
	if show {
		symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
		fmt.Println(RMB, symbol[RMB])
	}
}
