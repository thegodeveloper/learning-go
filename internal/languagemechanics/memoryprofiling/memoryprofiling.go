package memoryprofiling

import (
	"bytes"
	"io"
)

func Index(show bool) {

}

func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	// use a bytes buffer to provide a stream to process.
	input := bytes.NewReader(data)

	// the number of bytes we are looking for.
	size := len(find)

	// declare the buffers we need to process the stream.
	buf := make([]byte, size)
	end := size - 1

	// read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {
		// read in one byte from the input stream.
		if _, err := io.ReadFull(input, buf[end:]); err != nil {
			// flush the reset of the bytes we have.
			output.Write(buf[:end])
			return
		}

		// if we have a match, replace the bytes.
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			// read a new initial number of bytes.
			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}

			continue
		}

		// write the front byte since it has been compared.
		output.WriteByte(buf[0])

		// slice that front byte out.
		copy(buf, buf[1:])
	}
}

func assembleInputStream() []byte {
	return []byte("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis")
}

func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	// use a bytes buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// the number of bytes we are looking for.
	size := len(find)

	// declare the buffers we need to process the stream.
	//buf := make([]byte, size)
	buf := make([]byte, 5)

	end := size - 1

	// read in an initial number of bytes we need to get started.
	if n, err := input.Read(buf[:end]); err != nil || n < end {
		output.Write(buf[:n])
		return
	}

	for {
		// read in one byte from the input stream.
		if _, err := input.Read(buf[end:]); err != nil {
			// flush the reset of the bytes we have.
			output.Write(buf[:end])
			return
		}

		// if we have a match, replace the bytes.
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			// read a new initial number of bytes.
			if n, err := input.Read(buf[:end]); err != nil || n < end {
				output.Write(buf[:n])
				return
			}

			continue
		}

		// write the front byte since it has been compared.
		output.WriteByte(buf[0])

		// slice that front byte out
		copy(buf, buf[1:])
	}
}
