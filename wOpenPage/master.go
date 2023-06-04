package wOpenPage

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Open Web Page")
		fmt.Println(OnPage("https://www.dolar-colombia.com/"))
	}
}

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
