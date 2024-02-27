package wtest

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond also uses stuff set up in TestMain", testTime)
}

func Test_addNumber(t *testing.T) {
	result := addNumbers(4, 3)
	if result != 7 {
		t.Error("incorrect result: expected 7, got", result)
	}
}

func TestFileLen(t *testing.T) {
	result, err := FileLen("../README.md", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	if result != 655 {
		t.Error("Expected 655, got", result)
	}
}
