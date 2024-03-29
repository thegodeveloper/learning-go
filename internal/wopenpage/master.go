package wopenpage

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://www.dolar-colombia.com/"

func Master(show bool) {
	if show {
		fmt.Println("-- Open Web Page")
		fmt.Println(OnPage(URL))
	}
}

func OnPage(url string) string {
	res, err := http.Get(url)
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
