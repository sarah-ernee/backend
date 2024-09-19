package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type requestBody struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type responseBody struct {
	Message string      `json:"message"`
	Data    requestBody `json:"data"`
}

func createJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	resp := responseBody{
		Message: "JSON object created successfully",
		Data:    req,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	resp := responseBody{
		Message: fmt.Sprintf("JSON object with ID %s retrieved successfully", id),
		Data: requestBody{
			ID:          id,
			Title:       "Updated Title",
			Description: "Updated description",
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func updateJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	resp := responseBody{
		Message: fmt.Sprintf("JSON object with ID %s updated successfully", req.ID),
		Data:    req,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func deleteJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	resp := responseBody{
		Message: fmt.Sprintf("JSON object with ID %s deleted successfully", id),
		Data: requestBody{
			ID:          id,
			Title:       "Deleted Title",
			Description: "Deleted description",
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/items", createJSON).Methods("POST")
	r.HandleFunc("/items/{id}", getJSON).Methods("GET")
	r.HandleFunc("/items/{id}", updateJSON).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteJSON).Methods("DELETE")

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),        // Allow requests from this origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Allow these methods
		handlers.AllowedHeaders([]string{"Content-Type"}),                 // Allow these headers
	)

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", corsHandler(r)))
}
