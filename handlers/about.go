package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		PageTitle: "About",
	}

	err := h.Tmpl.ExecuteTemplate(w, "layout.html", data)
	fmt.Fprintln(w, "ABOUT PAGE HIT")

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
