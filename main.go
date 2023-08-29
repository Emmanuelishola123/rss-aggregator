package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Retrieve Port number from environment variables
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("Port number not specified")
	}

	// Create router
	router := chi.NewRouter()

	// Config cors
	router.Use(cors.Handler((cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})))

	// Create new router - v1
	router_v1 := chi.NewRouter()

	router_v1.Get("/ready", handlerReadiness)
	router_v1.Get("/err", handlerError)

	router.Mount("/v1", router_v1)

	// Create Server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	log.Printf("Server running on port %v", portNumber)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
