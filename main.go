package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("/contact", contactHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Portfolio running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
