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
