package handlers

import (
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "Home",
	}

	h.Tmpl.ExecuteTemplate(w, "layout.html", data)
}
