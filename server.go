package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/view/", makeHandler(loadHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	http.ListenAndServe(":8080", nil)
}
