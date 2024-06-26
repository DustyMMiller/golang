package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Starting webserver\n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
