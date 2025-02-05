package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, status, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{Status: status, Message: message}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding JSON:", err)
	}
}
