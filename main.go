package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}
