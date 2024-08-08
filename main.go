package main

import (
	"fmt"
	"funcs/funcs"
	"net/http"
)

func main() {
	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/ascii-art", funcs.AsciiArtHandler)
	// http.HandleFunc("/*", NotFoundHandler)
	fmt.Println("Your website is running at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "404 Not Found", http.StatusNotFound)
// }
