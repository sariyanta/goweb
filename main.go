package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, %q", r.URL.Path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Wrote %d bytes", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
