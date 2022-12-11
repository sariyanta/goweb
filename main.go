package main

import (
	"fmt"
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
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 3)
	fmt.Fprintf(w, "This is the about page. Sum is %d", sum)
}

func addValue(x, y int) int {
	return x + y
}
