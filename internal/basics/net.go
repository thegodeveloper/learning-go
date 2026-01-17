package basics

import (
	"fmt"
	"io"
	"net/http"
)

const staticUrl = "http://www.google.com"

// Net demonstrates using Go standard library to communicate with a URL
func Net(show bool) {
	if show {
		fmt.Println("Using the Go Standard Library to communicate with an URL")

		resp, err := http.Get(staticUrl)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}(resp.Body)

		fmt.Println("HTTP Response Status:", resp.Status)
	}
}
