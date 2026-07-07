package main

import (
	"go-web-server-test/handlers"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// we be routin' thangs, cuh!
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	tmpl := template.Must(template.ParseGlob("templates/*")) // dynamic handling of templates

	// i tried to add seperate "views" in static but it doesnt seem to wanna work
	// so i will just keep all of the pages in templates, but i dont think that it's
	// necessarily the cleanest way of doing this. oh welp
	h := &handlers.Handler{
		Tmpl: tmpl,
	}

	// handle functions here, they're located in "handlers". how simple, right?
	r.HandleFunc("/", h.Home)
	r.HandleFunc("/about", h.About)
	r.HandleFunc("/fun", h.Fun) // because we all need fun in our lives ;)

	http.ListenAndServe(":8000", r)
}
