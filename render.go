package main

import (
	"book/models"
	"net/http"
	"text/template"
)

func renderTemplate(option string, res http.ResponseWriter, p *models.Page) {
	error := templates.ExecuteTemplate(res, option+".html", p)
	if error != nil {
		http.Error(res, error.Error(), http.StatusInternalServerError)
	}
}

var templates = template.Must(template.ParseFiles("view.html", "edit.html"))
