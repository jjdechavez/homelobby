package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/jjdechavez/homelobby/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/health"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.HomeHandler)
	r.Get("/notes", handlers.NotesHandler)
	r.Get("/payments", handlers.PaymentsHandler)

	port := os.Getenv("PORT")
	fmt.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, r)
}
