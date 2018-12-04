package main

import (
	"log"
	"net/http"

	"wsf/Partiel-Devops/3A-Devops/handler"
)

func main() {

	myHandler := handler.NewHandler()

	server := &http.Server{
		Addr:    ":8080",
		Handler: myHandler,
	}

	log.Fatal(server.ListenAndServe())
}
