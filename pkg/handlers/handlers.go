package handlers

import (
	"net/http"

	"github.com/sariyanta/goweb/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page")
}
