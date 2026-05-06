package handlers

import (
	"net/http"
)

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "About",
	}

	h.Tmpl.ExecuteTemplate(w, "layout.html", data)
}
