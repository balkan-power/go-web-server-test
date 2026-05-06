package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Fun(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "Fun :)",
	}

	fmt.Printf("VIDEO: %q\n", data)

	err := h.Tmpl.ExecuteTemplate(w, "fun", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
