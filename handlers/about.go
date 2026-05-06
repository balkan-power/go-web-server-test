package handlers

import (
	"net/http"
)

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "About",
	}

	err := h.Tmpl.ExecuteTemplate(w, "about", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
