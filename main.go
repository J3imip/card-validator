package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/J3imip/card-validator/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Use(middleware.Logger)
	r.Post("/validate", handlers.ValidateCard)

	fmt.Println("Server running on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
