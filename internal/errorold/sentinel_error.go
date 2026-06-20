package errorold

import (
	"archive/zip"
	"bytes"
	"fmt"
)

// SentinelError demonstrates using sentinel errors
func SentinelError(show bool) {
	if show {
		data := []byte("This is not a zip file")
		notAZipFIle := bytes.NewReader(data)
		_, err := zip.NewReader(notAZipFIle, int64(len(data)))
		if err == zip.ErrFormat {
			fmt.Println("Told you so")
		}
	}
}
