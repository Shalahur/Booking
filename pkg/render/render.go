package render

import (
	"Booking/pkg/config"
	"Booking/pkg/models"
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var appConfig *config.AppConfig

func NewTemplates(configParam *config.AppConfig) {
	appConfig = configParam
}
func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

func RenderTemplate(writer http.ResponseWriter, request *http.Request, tmpl string, data *models.TemplateData) {
	var templateCache map[string]*template.Template

	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	templateObject, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	bufferObject := new(bytes.Buffer)

	data = AddDefaultData(data, request)

	err := templateObject.Execute(bufferObject, data)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

	_, err = bufferObject.WriteTo(writer)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache: create the template cache
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		//TODO: Funcs(functions) this is not clear to me
		templateStructure, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateStructure, err = templateStructure.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateStructure
	}
	return myCache, nil
}
