package funcs

import (
	"html/template"
	"net/http"
	"strings"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	asciiArt := r.URL.Query().Get("art")
	if asciiArt == "" {
		http.Error(w, "No ASCII art found", http.StatusNotFound)
		return
	}

	asciiArt = strings.ReplaceAll(asciiArt, "%0A", "\n")

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
