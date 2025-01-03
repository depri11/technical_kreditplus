package main

import (
	"log"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/server"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})

	r.Use(corsMiddleware.Handler)

	s := server.NewServer(r)

	log.Fatal(s.Run())
}
