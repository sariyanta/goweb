package main

import (
	"fmt"
	"net/http"

	"github.com/sariyanta/goweb/pkg/handlers"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("Server is running on port", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
