package wErrorOld

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func sentinelError(show bool) {
	if show {
		data := []byte("This is not a zip file")
		notAZipFIle := bytes.NewReader(data)
		_, err := zip.NewReader(notAZipFIle, int64(len(data)))
		if err == zip.ErrFormat {
			fmt.Println("Told you so")
		}
	}
}
