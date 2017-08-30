package main

import (
	"net/http"
	"html/template"
	"path"
	"encoding/json"
)

var templates = template.Must(template.ParseFiles(path.Join("resources", "react-app", "build", "index.html")))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/json", serveJson)
	fs1 := http.FileServer(http.Dir(path.Join("resources", "react-app", "build", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs1))
	http.ListenAndServe(":8080", nil)
}

func serveJson(w http.ResponseWriter, r *http.Request) {
	profile := SomeStruct{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, someStruct *SomeStruct) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", someStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type SomeStruct struct {
	Name    string
	Hobbies []string
}
