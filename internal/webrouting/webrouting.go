package webrouting

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(show bool) {
	if show {
		r := mux.NewRouter()

		r.HandleFunc("/books/{title}/title/{page}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			title := vars["title"]
			page := vars["page"]

			fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
		})

		http.ListenAndServe(":80", r)
	}
}
