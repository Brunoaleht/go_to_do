package handlers

import (
	"encoding/json"
	"log"
	"modules/models"
	"net/http"
	"strconv"
)

func Paginate(w http.ResponseWriter, r *http.Request) {
		// Extract query parameters
    query := r.URL.Query()
    title := query.Get("title")
    description := query.Get("description")
    completedStr := query.Get("completed")
		

		// Extract the completed query parameter
		completed, err := strconv.ParseBool(completedStr)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		//Create search params
		searchParams := models.TodoSearchParams{
			Title: title,
			Description: description,
			Completed: &completed,
		}

		//Get the Todos from the database
		todos, err := models.Paginate(searchParams)
		if err != nil {
			log.Printf("Error getting the Todos: %v", err)
		}

		//Send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
}