package main

import (
	"fmt"
	"log"
	"modules/configs"
	"modules/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Printf("Error LoadConfig: %v", err)
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/", handlers.Create)
	router.Patch("/{id}", handlers.Updated)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/{id}", handlers.Get)
	router.Get("/", handlers.Paginate)

	http.ListenAndServe(fmt.Sprintf(": %s", configs.GetServerPort()), router)

}