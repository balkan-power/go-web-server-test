package handlers

import (
	"net/http"
)

func (h *Handler) Fun(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "Fun :)",
	}

	err := h.Tmpl.ExecuteTemplate(w, "fun", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
