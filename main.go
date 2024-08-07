package main

import (
	"fmt"
	"funcs/funcs"
	"net/http"
)

func main() {
	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/ascii-art", funcs.AsciiArtHandler)
	fmt.Println("Your website is running at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
