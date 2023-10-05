package main

import (
	"book/models"
	"net/http"
	"regexp"
)

var valURL = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, t string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := valURL.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func loadHandler(res http.ResponseWriter, req *http.Request, title string) {

	p, e := models.LoadPage(title + ".txt")
	if e != nil {
		http.Redirect(res, req, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate("view", res, p)

}

func editHandler(res http.ResponseWriter, req *http.Request, title string) {

	p, e := models.LoadPage(title + ".txt")
	if e != nil {
		p = &models.Page{Title: title, Body: []byte("")}
	}
	renderTemplate("edit", res, p)

}

func saveHandler(res http.ResponseWriter, req *http.Request, title string) {

	p := &models.Page{Title: title, Body: []byte(req.FormValue("changed"))}
	e := p.SavePage()
	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
	}
	http.Redirect(res, req, "/view/"+p.Title, http.StatusFound)
}
