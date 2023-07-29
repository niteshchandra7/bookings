package render

import (
	"net/http"
	"testing"

	"github.com/niteshchandra7/bookings/internals/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error("failed")
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplate()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-session"))

	r = r.WithContext(ctx)
	return r, nil

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplate()
	if err != nil {
		t.Error("error while creating template cache")
	}
}
