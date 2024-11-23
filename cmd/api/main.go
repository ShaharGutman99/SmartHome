package main

import (
	"log"
	"net/http"
	"github.com/ShaharGutmanCoding/SmartHome/internal/routes"
)

func main() {
	// Initialize the router
	router := routes.SetupRoutes()

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
