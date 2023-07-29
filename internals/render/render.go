package render

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/niteshchandra7/bookings/internals/config"
	"github.com/niteshchandra7/bookings/internals/models"
)

var tc = make(map[string]*template.Template)
var app *config.AppConfig
var pathToTemplates string = "./templates"

var Funcs = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	//var ok bool
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	//td.Error, ok = app.Session.Get(r.Context(), "error").(string)
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplate()
		if err != nil {
			log.Fatal(err)
			return errors.New("can't get template from cache")
		}
	}
	_, isFound := tc[tmpl]
	if !isFound {
		log.Fatal("could not get template from template cache")
		return errors.New("can't get template from cache")
	}
	td = AddDefaultData(td, r)
	err = tc[tmpl].Execute(w, td)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return nil

}

func CreateTemplate() (map[string]*template.Template, error) {
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return tc, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(Funcs).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			log.Fatal("layout file not found")
			return nil, err
		}
		t, err = t.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return tc, err
		}
		tc[name] = t
	}
	return tc, nil
}
