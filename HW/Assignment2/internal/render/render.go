package render

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
)

var functions = template.FuncMap{}
var pathToTemplates = "./templates"

func CreateTemplateCahce() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
	}
	matches

}
