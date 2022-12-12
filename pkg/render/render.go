package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// map of template cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see template is in cache
	_, inCache := tc[t]

	if !inCache {
		// Create template cache
		log.Println("Creating template cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// get template from cache
		log.Println("Getting template from cache")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// Add template to cache
	tc[t] = tmpl

	return nil
}
