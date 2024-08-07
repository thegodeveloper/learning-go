package arrays

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func symbols(show bool) {
	if show {
		symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
		fmt.Println(RMB, symbol[RMB])
	}
}
