package handlers

import (
	"math/rand"
	"net/http"
)

func (h *Handler) Fun(w http.ResponseWriter, r *http.Request) {

	// define what quotes are used, HERE!
	quotes := []string{
		// the quotes in dutch, i liked them so i added them here
		"moet gwn goed functionere",
		"zeg ik je nie, ik ben geen verraaier",
		"zo zijn ze braaf",
		"ik zeg er niks van",
		"ik kan niet leven zonder internet",
		"hey, da's 't ibahesj",
		"mensen hebben mij dat al meer dan 20 keer genoemd",
		"we hebben een serieus probleem",
		"zeg ik je nie",
		"je moeder",
		"hier werken ze met eentjes en nulletjes",
		"dat zeg ik niet tegen jou, ik ben geen verrader",
		"dat zeg ik niet tegen jou",
		"IBAHESJ",
		"HEY PIKKIE",
		"ik praat tegen mezelf",
		"ja ja ja ja ja ja ja ja ja ja ja ja ja ja ja ja ja",

		// the quotes in english, from wikiquote, of Richard Stallman.
		"Every decision a person makes stems from the person's values and goals.",
		"Standing up to an evil system is exhilarating, and now I have a taste for it.",
		"I've always lived cheaply. I live like a student, basically. And I like that, because it means that money is not telling me what to do. I can do what I think is important for me to do. It freed me to do what seemed worth doing.",
		"Value your freedom or you will lose it, teaches history. 'Don't bother us with politics,' respond those who don't want to learn. ",
		"Odious ideas are not entitled to hide from criticism behind the human shield of their believers' feelings.",
		"Geeks like to think that they can ignore politics, you can leave politics alone, but politics won't leave you alone.",
		"People said I should accept the world. Bullshit! I don't accept the world.",
		"I don't have a problem with someone using their talents to become successful, I just don't think the highest calling is success. Things like freedom and the expansion of knowledge are beyond success, beyond the personal. Personal success is not wrong, but it is limited in importance, and once you have enough of it it is a shame to keep striving for that, instead of for truth, beauty, or justice.",
		"While corporations dominate society and write the laws, each advance in technology is an opening for them to further restrict its users.",
	}

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
	}

	random_quote := quotes[rand.Intn(len(quotes))]
	random_meme := memes[rand.Intn(len(memes))]

	data := PageData{
		PageTitle: "Fun :)",
		Quote:     random_quote,
		Meme:      random_meme,
	}

	err := h.Tmpl.ExecuteTemplate(w, "fun", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
