package funcs

import (
	"html/template"
	"net/http"
	"net/url"
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

		if bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy" || len(text) > 1000 {
			http.Error(w, "Bad request", http.StatusNotFound)
			return
		}

		banner, err := GetBanner(bannerName)
		if err != nil {
			http.Error(w, "Bad request", http.StatusInternalServerError)
			return
		}

		asciiArt := GenerateASCIIArt(text, banner)

		// Safely encode ASCII art for URL query
		encodedArt := url.QueryEscape(asciiArt)
		http.Redirect(w, r, "/ascii-art?art="+encodedArt, http.StatusSeeOther)
		return
	}

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
