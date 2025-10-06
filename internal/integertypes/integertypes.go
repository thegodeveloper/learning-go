package integertypes

import "fmt"

func Index(show bool) {
	if show {
		fmt.Println("The Integer Types in Go.")

		int8Type(false)
		int16Type(false)
		int32Type(false)
		int64Type(false)
		intType(false)
		uint8Type(false)
		byteType(false)
		uint16Type(false)
		uint32Type(false)
		uint64Type(false)
	}
}

func int8Type(show bool) {
	if show {
		fmt.Println("int8 Type")
		var i int8 = 127
		fmt.Println("i positive value: ", i)
		i = -127
		fmt.Println("i negative value: ", i)
	}
}

func int16Type(show bool) {
	if show {
		fmt.Println("int16 Type")
		var i int16 = 32767
		fmt.Println("i positive value: ", i)
		i = -32768
		fmt.Println("i negative value: ", i)
	}
}

func int32Type(show bool) {
	if show {
		fmt.Println("int32 Type")
		var i int32 = 2147483647
		fmt.Println("i positive value: ", i)
		i = -2147483648
		fmt.Println("i negative value: ", i)
	}
}

func int64Type(show bool) {
	if show {
		fmt.Println("int64 Type")
		var i int64 = 9223372036854775807
		fmt.Println("i positive value: ", i)
		i = -9223372036854775808
		fmt.Println("i negative value: ", i)
	}
}

func intType(show bool) {
	if show {
		fmt.Println("int64 Type")
		var i int = 9223372036854775807
		fmt.Println("i positive value: ", i)
		i = -9223372036854775808
		fmt.Println("i negative value: ", i)
	}
}

func uint8Type(show bool) {
	if show {
		fmt.Println("uint8 Type")
		var i uint8 = 0
		fmt.Println("i min value: ", i)
		i = 255
		fmt.Println("i max value: ", i)
	}
}

func byteType(show bool) {
	if show {
		fmt.Println("uint8 Type")
		var i byte = 0
		fmt.Println("i min value: ", i)
		i = 255
		fmt.Println("i max value: ", i)
	}
}

func uint16Type(show bool) {
	if show {
		fmt.Println("uint16 Type")
		var i uint16 = 0
		fmt.Println("i min value: ", i)
		i = 65535
		fmt.Println("i max value: ", i)
	}
}

func uint32Type(show bool) {
	if show {
		fmt.Println("uint32 Type")
		var i uint32 = 0
		fmt.Println("i min value: ", 0)
		i = 4294967295
		fmt.Println("i max value: ", i)
	}
}

func uint64Type(show bool) {
	if show {
		fmt.Println("uint64 Type")
		var i uint64 = 0
		fmt.Println("i min value: ", i)
		i = 18446744073709551615
		fmt.Println("i max value: ", i)
	}
}
