package slices

import (
	"fmt"
	"os"
)

func SlicesAsBuffers(show bool, fileName string) error {
	if show {
		fmt.Println("--- Slice as Buffers")

		f, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer f.Close()

		data := make([]byte, 100)
		for {
			count, err := f.Read(data)
			if err != nil {
				return err
			}
			if count == 0 {
				return nil
			}
			fmt.Printf("Data: %v\n", data[:count])
		}
	} else {
		return nil
	}
}
