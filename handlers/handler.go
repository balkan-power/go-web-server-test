package handlers

import (
	"html/template"
	"math/rand"
	"net/http"
)

// set structures here for the handler so the pages can process the... you know what i mean

type Handler struct {
	Tmpl *template.Template
}

func randomMemeHandler(w http.ResponseWriter, r *http.Request) {
	memes := []string{
		"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		"https://www.youtube.com/watch?v=TNUXzTP79e0",
		"https://www.youtube.com/watch?v=3c2wdlxLy7Q",
		"https://www.youtube.com/watch?v=XcyzLyZeqf4?t=43",
		"https://www.youtube.com/watch?v=j308rH1j_bg",
		"https://www.youtube.com/watch?v=vbIbIurceNU",
		"https://www.youtube.com/watch?v=EAzskL3gNxc",
		"https://www.youtube.com/watch?v=dQa8lydtFmE",
		"https://www.youtube.com/watch?v=s0Gpd2ooB9w",
		"https://www.youtube.com/watch?v=wa5inGuht_o",
		"https://www.youtube.com/watch?v=bPzVV_5sQtc",
		"https://www.youtube.com/watch?v=sUSN7fqVBio",
	}

	url := memes[rand.Intn(len(memes))]
	http.Redirect(w, r, url, http.StatusFound)
	http.HandleFunc("/random-meme", randomMemeHandler)
}

// very important info :)
type PageData struct {
	PageTitle string
	Quote     string
	Meme      string
}
