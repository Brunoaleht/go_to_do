package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"modules/models"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	// Decode the incoming Todo json
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
    log.Printf("Error decoding the incoming Todo JSON: %v", err)
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
	}

	// Insert the Todo into the database
	id, err := models.Insert(todo)

	var response map[string]any
	StatusCode := http.StatusOK

	if err != nil {
		StatusCode = http.StatusInternalServerError
		response = map[string]any{
			"Error": true, 
			"Message": fmt.Sprintf("Failed to insert the Todo: %v", err), 
		}
	} else {
		response = map[string]any{
			"Error": false, 
			"Message": fmt.Sprintf("Todo inserted successfully ID: %v", id), 
			"Id": id, 
		}
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(response)

}