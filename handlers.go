package main

import "net/http"

// Project describes one card on the Projects page.
type Project struct {
	Name        string
	Description string
	URL         string
}

// PageData is what every template receives.
type PageData struct {
	Title    string
	Projects []Project
}

var projects = []Project{
	{
		Name:        "ASCII-Art-web",
		Description: "A Go web app: type text, pick shadow, standard or thinkertoy, and get it back as ASCII art.",
		URL:         "https://github.com/ejembiedwin36-oss/ASCII-Art-web",
	},
	{
		Name:        "Go-Reload",
		Description: "A modular Go text-processing engine: hex and binary conversion, case changes, a/an fixes, punctuation and quote formatting.",
		URL:         "https://github.com/ejembiedwin36-oss/Go-Reload",
	},
	{
		Name:        "ASCII-Art",
		Description: "ASCII Art Generator takes a string as a command-line argument and outputs it as a graphic representation using ASCII characters. It supports three banner styles and handles special characters, numbers, spaces, and newlines.",
		URL:         "https://github.com/ejembiedwin36-oss/ASCII-Art",
	},
	{
		Name:        "the-codecrafters",
		Description: ".Yeah the codecrafteers is mini-project where i learn and build CLI Calculator that work with terminal, and w3school is were i learn the basics for the project, then the base converter is is just basically how to convert numbers like converting Hexadecimal to binary and from binary to Decimal.",
		URL:         "https://github.com/ejembiedwin36-oss/the-codecrafters",
	},
}

// 1. Home: the "/" route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	renderPage(w, "home.html", PageData{Title: "Home"})
}

// 2. About
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	renderPage(w, "about.html", PageData{Title: "About"})
}

// 3. Projects
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	renderPage(w, "projects.html", PageData{Title: "Projects", Projects: projects})
}

// 4. Contact
func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	renderPage(w, "contact.html", PageData{Title: "Contact"})
}