package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError retuens error if form of application/json format
func respondWithError(w http.ResponseWriter, statusCode int, msg string) {

	if statusCode > 499 {
		log.Println("Responding with 5xx response error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errResponse{
		Error: msg,
	})

}

// respondWithJSON retuens any http routes with application/json format
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
