package funcs

import (
	"html/template"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	asciiArt := r.FormValue("art")
	if asciiArt == "" {
		http.Error(w, "No ASCII art found", http.StatusBadRequest)
		return
	}

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
