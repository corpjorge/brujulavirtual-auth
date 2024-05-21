package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := errorResponse{
		Error: errorMessage,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Default().Printf("Error encoding the error response: %v", err)
	}

	log.Printf("Error: %v", errorResponse)
}
