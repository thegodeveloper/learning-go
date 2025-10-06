package basics

import (
	"fmt"
	"io"
	"net/http"
)

func net(show bool) {
	if show {
		fmt.Println("Go Standard Library")

		resp, err := http.Get("http://www.google.com")
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
