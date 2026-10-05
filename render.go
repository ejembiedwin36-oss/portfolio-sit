package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

// renderPage loads layout.html plus one page file, fills in the data,
// and writes the finished HTML to the browser.
func renderPage(w http.ResponseWriter, page string, data PageData) {
	t, err := template.ParseFiles("templates/layout.html", "templates/"+page)
	if err != nil {
		log.Println("template load error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Build the page in memory first, so a failure never sends half a page.
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "layout", data); err != nil {
		log.Println("template run error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
