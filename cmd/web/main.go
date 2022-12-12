package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sariyanta/goweb/pkg/config"
	"github.com/sariyanta/goweb/pkg/handlers"
	"github.com/sariyanta/goweb/pkg/render"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
