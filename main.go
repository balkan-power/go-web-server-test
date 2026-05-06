package main

import (
	"go-web-server-test/handlers"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	tmpl := template.Must(template.ParseGlob("templates/*")) // dynamic handling of templates

	h := &handlers.Handler{
		Tmpl: tmpl,
	}

	// handle functions here, they're located in "handlers". how simple, right?
	r.HandleFunc("/", h.Home)
	r.HandleFunc("/about", h.About)
	r.HandleFunc("/fun", h.Fun) // because we all need fun in our lives ;)

	http.ListenAndServe(":8000", r)
}
