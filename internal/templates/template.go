package templates

import (
	"errors"
	"html/template"
	"log"
	"path/filepath"

	"github.com/KoLLlaka/libraryAPI/internal/model"
)

type PageData struct {
	Title string
	Books []*model.Book
}

type TemplateCashe struct {
	tmplCashe map[string]*template.Template
}

func NewTemplateCashe() (*TemplateCashe, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./html/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	// tmpl := template.Must(template.ParseFiles("./html/index.go.html"))

	// mapTemp.tmplCashe["index"] = tmpl

	return &TemplateCashe{tmplCashe: cache}, nil
}

func (t *TemplateCashe) Render(name string) *template.Template {
	tmpl, ok := t.tmplCashe[name]
	if !ok {
		log.Fatal(errors.New("Template is not exist"))
		return nil
	}

	return tmpl
}
