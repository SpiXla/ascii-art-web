package funcs

import (
	"html/template"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		text := r.FormValue("text")
		bannerName := r.FormValue("banner")

		banner, err := GetBanner(bannerName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		asciiArt := GenerateASCIIArt(text, banner)

		// Redirect with ASCII art as a query parameter
		http.Redirect(w, r, "/ascii-art?art="+strings.ReplaceAll(asciiArt, "\n", "%0A"), http.StatusSeeOther)
	} else {
		tmpl, err := template.ParseFiles("html/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
