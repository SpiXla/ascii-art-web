package funcs

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Redirect with ASCII art as a query parameter
		// http.Redirect(w, r, "/ascii-art?art="+strings.ReplaceAll(asciiArt, "\n", "%0A"), http.StatusSeeOther)
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
