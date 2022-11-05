package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := []byte("nirvana")
	c2 := []byte("foo fighters")
	fmt.Printf("c1 value: %v\n", c1)
	fmt.Printf("c2 value: %v\n", c2)
	fmt.Printf("Number of different bits in two SHA256 hashes: %d\n", sha256PopCount(c1, c2))
}

func sha256PopCount(a, b []byte) int {
	digesta := sha256.Sum256(a)
	digestb := sha256.Sum256(b)
	return popCount(digesta, digestb)
}

func popCount(a, b [32]byte) int {
	pop := 0
	for i := range a {
		pop += int(pc[a[i]^b[i]])
	}
	return pop
}
