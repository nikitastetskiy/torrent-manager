package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitastetskiy/torrent-api/internal/handler"
)

func main() {
	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
