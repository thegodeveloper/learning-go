package wrand

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Working with Random Values")
		definition()
	}
}

func definition() {
	s := seedRand()
	fmt.Println("value of random s:", s)
}

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}
