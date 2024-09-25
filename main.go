package main

import (
	"fmt"
	"go-web-boilerplate-v1/config"
	"go-web-boilerplate-v1/middleware"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, I'm flash :)")
}

func main() {
	// log
	filename := "./test.log" // todo config
	logger, err := middleware.GetLogger(filename)
	if err != nil {
		fmt.Println("getLogger err")
	}
	logger.Info("111111111")

	// config
	// env := "aks" // todo Dockerfile
	env := "eks" // todo Dockerfile
	if err := config.GetConfig(env); err != nil {
		log.Fatal("Cannot get config:", err)
	}
	name := viper.GetString("name")
	logger.Info(name)

	// web
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
