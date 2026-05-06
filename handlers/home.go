package handlers

import (
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "Home",
	}

	err := h.Tmpl.ExecuteTemplate(w, "home", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
