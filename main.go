package main

import (
	"fmt"
	"go-web-server-test/handlers"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	tmpl := template.Must(template.ParseGlob("templates/*.html")) // dynamic handling of templates

	h := &handlers.Handler{
		Tmpl: tmpl,
	}

	// handle functions here, they're located in "handlers". how simple, right?
	r.HandleFunc("/", h.Home)
	r.HandleFunc("/about", h.About)
	r.HandleFunc("/fun", h.Fun) // because we all need fun in our lives ;)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Site reached!")
	})

	http.ListenAndServe(":80", r)
}
