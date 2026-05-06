package handlers

import (
	"math/rand"
	"net/http"
)

func (h *Handler) Fun(w http.ResponseWriter, r *http.Request) {

	// define what video IDs are used, HERE!
	videos := []string{
		"dQw4w9WgXcQ",
		"9bZkp7q19f0",
		"kxopViU98Xo",
	}

	random_vid := videos[rand.Intn(len(videos))]

	data := PageData{
		PageTitle: "Fun :)",
		VideoID:   random_vid,
	}

	err := h.Tmpl.ExecuteTemplate(w, "fun", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
