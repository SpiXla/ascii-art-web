package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("Your website is running at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		text := r.FormValue("text")
		bannerName := r.FormValue("banner")

		banner, err := getBanner(bannerName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		asciiArt := generateASCIIArt(text, banner)

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

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
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

func getBanner(bannerName string) ([]string, error) {
	// Read banner from file
	bannerData, err := os.ReadFile("banners/" + bannerName + ".txt")
	if err != nil {
		return nil, err
	}

	// Split banner data into lines
	banner := strings.Split(string(bannerData), "\n")

	return banner, nil
}

func generateASCIIArt(text string, banner []string) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					// Or any other placeholder
					continue
				}
				index := 9*(int(r)-32) + i + 1
				result.WriteString(banner[index])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
