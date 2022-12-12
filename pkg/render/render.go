package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//  get the template cache from the app config

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)

	err = t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	// loop through the pages
	for _, page := range pages {
		// get the file name
		name := filepath.Base(page)

		// parse the template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		// add to template cache
		templateCache[name] = ts
	}

	return templateCache, nil
}
