package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wsf/Partiel-Devops/3A-Devops/handler"
)

func main() {

	connection := os.Getenv("DATABASE_URL")

	fmt.Println(connection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	myHandler := handler.NewHandler()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: myHandler,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running at", port)

	log.Fatal(server.ListenAndServe())
}
