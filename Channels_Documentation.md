
# Go Channels: Concepts and Examples

## Overview

This document provides an explanation of the Go channels examples found in the `internal/channels` package. Go channels are a powerful feature for communication and synchronization between goroutines. This package demonstrates various aspects of channel usage, from basic definitions to more advanced concepts.

## Core Concepts

The main entry point for running all the channel demonstrations is the `definition()` function, which is called by the exported `Definition(bool)` function. The `Run(bool)` function is the main entry point for the TUI.

### `definition.go`

This file orchestrates the demonstration of several key channel operations in a sequence.

```go
package channels

import "fmt"

func definition() {
	fmt.Println("-- Buffered Channel")
	bufferedChannel()

	fmt.Println("-- Unbuffered Channel")
	unbufferedChannel()

	fmt.Println("-- Read and Write from/to the Channel")
	readAndWriteToFromChannel()

	fmt.Println("-- Closing a Channel")
	closedChannel()
}
```

### `channels.go`

This file contains the main `Run` function that is used by the TUI to execute the examples in this package.

```go
package channels

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Channels")
		definition()

		fmt.Println("-- Channels as Function Parameters")
		channelsAsFunctionParameters()
	}
}

// Definition demonstrates all basic channel operations
func Definition(show bool) {
	if !show {
		return
	}
	definition()
}
```

## Channel Types

### Unbuffered Channels (`unbufferedChannel.go`)

An unbuffered channel has no capacity. A send operation on an unbuffered channel blocks until a receiver is ready to receive the value. This ensures synchronization between the sender and receiver.

```go
package channels

import "fmt"

// UnbufferedChannel demonstrates unbuffered channel usage
func UnbufferedChannel(show bool) {
	if !show {
		return
	}
	unbufferedChannel()
}

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		ch <- 77
	}()
	val := <-ch
	fmt.Println("value received from unbuffered channel val:", val)
}
```

### Buffered Channels (`bufferedChannel.go`)

A buffered channel has a specific capacity. A send operation on a buffered channel only blocks if the buffer is full. A receive operation blocks if the buffer is empty.

```go
package channels

import "fmt"

// BufferedChannel demonstrates buffered channel usage
func BufferedChannel(show bool) {
	if !show {
		return
	}
	bufferedChannel()
}

func bufferedChannel() {
	ch := make(chan int, 1)
	ch <- 77
	val := <-ch
	fmt.Println("value received from buffered channel val:", val)
}
```

## Channel Operations

### Reading from and Writing to a Channel (`readAndWriteToFromChannel.go`)

This example demonstrates the fundamental operations of sending a value to a channel in one goroutine and receiving it in another.

```go
package channels

import "fmt"

// ReadAndWriteToFromChannel demonstrates reading and writing to channels
func ReadAndWriteToFromChannel(show bool) {
	if !show {
		return
	}
	readAndWriteToFromChannel()
}

func readAndWriteToFromChannel() {
	ch := make(chan int)

	// write to the channel
	go func() {
		ch <- 77
	}()

	// read from the channel
	val := <-ch
	fmt.Println("value read from the channel:", val)

	// write to the channel again
	go func() {
		ch <- 777
	}()

	// read from the channel again
	val = <-ch
	fmt.Println("value read from the channel again:", val)
}
```

### Closing a Channel (`closedChannel.go`)

Channels can be closed to indicate that no more values will be sent on them. This is useful for communicating completion to the receiver. The receiver can test whether a channel has been closed by using a two-variable assignment: `val, ok := <-ch`. If `ok` is `false`, the channel is closed and `val` is the zero value for the channel's type.

```go
package channels

import "fmt"

// ClosedChannel demonstrates closing channels and detecting closed state
func ClosedChannel(show bool) {
	if !show {
		return
	}
	closedChannel()
}

func closedChannel() {
	ch := make(chan int, 7)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	close(ch)

	for {
		val, ok := <-ch
		if !ok { // if the channel is closed, ok will be false
			fmt.Println("channel closed!")
			break
		}
		fmt.Println("receiving from channel:", val)
	}
}
```

### Channels as Function Parameters (`channels_as_function_parameters.go`)

Channels can be passed to functions just like any other value. You can also specify the direction of a channel parameter, i.e., whether it's for sending (`chan<-`) or receiving (`<-chan`). This enforces type safety at compile time.

```go
package channels

import "fmt"

func f1(c chan int, x int) {
	fmt.Println("value of x:", x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println("value of x:", x)
	c <- x
}

// ChannelsAsFunctionParameters demonstrates using channels as function parameters
func ChannelsAsFunctionParameters(show bool) {
	if !show {
		return
	}
	channelsAsFunctionParameters()
}

func channelsAsFunctionParameters() {
	c1 := make(chan int)
	c2 := make(chan int)

	go f1(c1, 1)
	go f2(c2, 2)

	x := <-c1
	y := <-c2

	fmt.Println("value received", x, y)
}
```
