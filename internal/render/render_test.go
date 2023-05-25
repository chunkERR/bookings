package render

import (
	"net/http"
	"testing"

	"github.com/chunkERR/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	templateCache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.MyCache = templateCache
	
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(ww, r, "home.page.tmpl", &models.TemplateData{}) {

	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	context := r.Context()
	context, _ = session.Load(context, r.Header.Get("X-Session"))
	r = r.WithContext(context)

	return r, nil

}
