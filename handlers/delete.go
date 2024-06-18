package handlers

import (
	"encoding/json"
	"log"
	"modules/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// Parse the URL parameter
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting the URL parameter to int: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Delete the Todo from the database
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error deleting the Todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Validate the correct number of rows deleted
	if rows > 1 {
		log.Printf("Error deleting the Todo: %v", rows)
	}

	// Send the response
	resp := map[string]any{
		"Error":   false,
		"Message": "Todo deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}