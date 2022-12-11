package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	sum, err := divideValue(10.0, 1.0)

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	fmt.Fprintf(w, "Sum is %f", sum)
}

func addValue(x, y int) int {
	return x + y
}

func divideValue(x, y float32) (float32, error) {

	if y <= 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := x / y
	return result, nil
}
