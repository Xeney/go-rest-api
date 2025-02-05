package main

import (
	"http-server/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/greet", handlers.GreetHandler)
	http.HandleFunc("/create", handlers.CreateHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/update", handlers.UpdateHandler)
	log.Println("start server to http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
