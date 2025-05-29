package main

import (
	"fmt"
	"log"
	"os"

	"go_login/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Get server configuration from environment
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0" // default host
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	r := routes.SetupRouter()
	serverAddr := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Server starting on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
