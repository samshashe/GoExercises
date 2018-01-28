package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/home.html"))

func RenderTemplate(w http.ResponseWriter, data interface{}) {
	err := templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
