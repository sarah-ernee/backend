package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	log.Println("createJSON called")
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
	log.Println("getJSON called")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	resp := responseBody{
		Message: fmt.Sprintf("JSON object with ID %s retrieved successfully", id),
		Data: requestBody{
			ID:          id,
			Title:       "Lorem",
			Description: "The sly brown fox jumped over the fence.",
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func updateJSON(w http.ResponseWriter, r *http.Request) {
	log.Println("updateJSON called")
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
	log.Println("deleteJSON called")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	resp := responseBody{
		Message: fmt.Sprintf("JSON object with ID %s deleted successfully", id),
		Data: requestBody{
			ID:          id,
			Title:       "Deleted Title",
			Description: "Deleted description.",
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(loggerMiddleware)

	r.HandleFunc("/items", createJSON).Methods("POST")
	r.HandleFunc("/items/{id}", getJSON).Methods("GET")
	r.HandleFunc("/items", updateJSON).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteJSON).Methods("DELETE")

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
