package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, I'm flash :)")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
