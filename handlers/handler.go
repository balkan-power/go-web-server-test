package handlers

import (
	"html/template"
)

// set structures here for the handler so the pages can process the... you know what i mean

type Handler struct {
	Tmpl *template.Template
}

// very important info :)
type PageData struct {
	PageTitle string
	Quote     string
}
