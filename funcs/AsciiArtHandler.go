package funcs

import (
	"html/template"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	if len(text) > 100 || text == "" || (bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy") {
		http.Error(w, "Bad request 400", http.StatusBadRequest)
		return
	}
	banner, err := GetBanner(bannerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// asciiArt := r.FormValue("art")
	asciiArt := GenerateASCIIArt(text, banner)
	// asciiArt := r.URL.Query().Get("art")
	if asciiArt == "" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	// asciiArt = strings.ReplaceAll(asciiArt, "%0A", "\n")

	data := map[string]string{
		"ASCIIArt": asciiArt,
	}

	tmpl, err := template.ParseFiles("html/ascii_art.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
