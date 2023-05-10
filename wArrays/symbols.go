package wArrays

import "fmt"

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func symbols() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
}
