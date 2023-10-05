package models

import (
	"os"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) SavePage() error {
	return os.WriteFile(("pages/" + p.Title + ".txt"), p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	b, e := os.ReadFile("pages/" + title)
	if e != nil {
		return nil, e
	}
	return &Page{Title: strings.Split(title, ".")[0], Body: b}, nil
}
