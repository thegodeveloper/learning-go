package wTest

import (
	"fmt"
	"testing"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Tests")
	}
}

func addNumber(x, y int) int {
	return x + y
}

func Test_addNumber(t *testing.T) {
	result := addNumber(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
