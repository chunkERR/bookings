package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/chunkERR/bookings/pkg/config"
	"github.com/chunkERR/bookings/pkg/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) error {

	var templatecache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templatecache = app.MyCache

	} else {
		templatecache, _ = CreateTemplateCache()
	}

	t, ok := templatecache[tmpl]
	if !ok {
		log.Printf("template not found: %s\n", tmpl)
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, r)

	err := t.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		fmt.Println("Error writing template to browser", err)
	}
return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("error parsing template:", err)
			continue
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("error parsing template:", err)
			continue
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("error parsing template:", err)
				continue
			}

		}

		myCache[name] = ts
	}

	return myCache, nil
}
