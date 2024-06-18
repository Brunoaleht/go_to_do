package handlers

import (
	"encoding/json"
	"log"
	"modules/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Updated(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting the URL parameter to int: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//Parse the incoming Todo JSON
	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding the incoming Todo JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Update the Todo in the database
	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error updating the Todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//validation is correct number of rows updated
	if rows > 1 {
		log.Printf("Error updating the Todo: %v", rows)
	}

	//Send the response
	resp := map[string]any{
		"Error":   false,
		"Message": "Todo updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}