package main

import (
	"fmt"
	"go-web-boilerplate-v1/middleware"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, I'm flash :)")
}

func main() {
	filename := "./test.log"
	logger, err := middleware.GetLogger(filename)
	if err != nil {
		fmt.Println("getLogger err")
	}
	logger.Info("111111111")

	// web
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
